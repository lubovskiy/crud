package main

import (
	"context"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lubovskiy/crud/helpers/shutdown"
	"github.com/lubovskiy/crud/infrastructure/database"
	"github.com/lubovskiy/crud/internal/config"
	"github.com/lubovskiy/crud/internal/repository/phonebook"
	"github.com/lubovskiy/crud/internal/service"
	"github.com/lubovskiy/crud/pkg/crud"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

const (
	HTTPServerReadTimeout  = time.Minute
	HTTPServerWriteTimeout = time.Minute
)

func main() {
	ctx := context.Background()

	connAddress := config.GetConfigConnection()
	conn, err := database.NewConn(connAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	mux := http.NewServeMux()
	httpServer := &http.Server{
		Handler:      mux,
		ReadTimeout:  HTTPServerReadTimeout,
		WriteTimeout: HTTPServerWriteTimeout,
	}

	run(ctx, srv, httpServer, conn)
}

func run(ctx context.Context, grpcServer *grpc.Server, httpServer *http.Server, conn *pgxpool.Pool) {
	pb := phonebook.NewRepository(conn)

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	sigHandler := shutdown.TermSignalTrap()

	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	crud.RegisterContactsServer(grpcServer, service.NewPhonebookService(pb))

	go func() {
		err := grpcServer.Serve(ls)
		if err != nil && err != cmux.ErrServerClosed {
			log.Fatal("grpc server serve error")
		}
	}()

	mux := cmux.New(ls)
	httpListener := mux.Match(cmux.HTTP1Fast())

	go func() {
		err := httpServer.Serve(httpListener)
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("https server serve error")
		}
		logger.Info("http listener started")
	}()

	err = sigHandler.Wait(ctx)
	if err != nil && err != shutdown.ErrTermSig && err != context.Canceled {
		logger.Error("failed to caught signal")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	err = httpServer.Shutdown(ctx)
	if err != nil {
		logger.Error("failed to shutdown http server")
	}

	grpcServer.GracefulStop()
}
