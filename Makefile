ifeq ($(OS),Windows_NT)
	CUR_DIR=$(shell echo %CD%)
else
	CUR_DIR=$(shell pwd)
	SHELL:=/bin/bash
	ifeq ($(shell test -e ./deployment/.env && echo -n yes),yes)
		include ./deployment/.env
		export
	endif
endif

#K8S_VERSION=1.18.0
#KUBECTL_EXEC=minikube kubectl --

PROJECT=github.com/lubovskiy/crud
IMAGE=crud
TAG=latest

HELM_DIR=${CUR_DIR}/deployment/helm
RELEASE_NAME=crud

DC_FILES=-f ${CUR_DIR}/deployment/docker-compose.yml

.PHONY: vendor test.unit test.integration compile deploy delete logs.service gen.proto migrate \
deploy-pg delete-pg deploy-debug

vendor: ## Установка зависимостей
	go mod download

test.unit: ## Прогон юнит-тестов
	go test -v -count=1 -race -coverprofile=c.out ./...

test.integration: delete deploy-pg migrate ## Прогон интеграционных тестов
	PG_HOST=$(PG_EXT_HOST) PG_PORT=$(PG_EXT_PORT) PG_DEBUG= go test -v -count=1 -tags "integration" -race -coverprofile=c.out ./...
	go tool cover -html=c.out -o coverage.html

compile: ## Сборка локального образа
	docker build --no-cache -t ${IMAGE}:${TAG} --target build .

deploy: ## Локальный деплой
	cd deployment && docker-compose ${DC_FILES} -p ${RELEASE_NAME} up -d

delete: ## Удаление локального деплоя
	cd deployment && docker-compose ${DC_FILES} -p ${RELEASE_NAME} rm -sf

logs.service: ## Логи контейнера
	cd deployment && docker-compose ${DC_FILES} -p ${RELEASE_NAME} logs service

gen.proto: ## Генерация proto
	docker run -it --rm -v ${CUR_DIR}/api:/go/src/${PROJECT}/api -v ${CUR_DIR}/pkg:/go/src/${PROJECT}/pkg -v ${CUR_DIR}/bin:/go/src/${PROJECT}/bin ${IMAGE}:${TAG} bash ./scripts/generate.sh /go/src/${PROJECT}

migrate: ## Запуск миграций
	cd deployment && docker-compose ${DC_FILES} -p ${RELEASE_NAME} up migrate

deploy-pg: ## Локальный деплой базы данных
	cd deployment && docker-compose ${DC_FILES} -p ${RELEASE_NAME} up -d postgres

delete-pg: ## Удаление локального деплоя базы данных
	cd deployment && docker-compose ${DC_FILES} -p ${RELEASE_NAME} rm -sf postgres

deploy-debug: deploy-rmq deploy-pg migrate ## Локальный деплой для дебага

#.PHONY: env.create
#env.create:
#	$(KUBECTL_EXEC) create configmap ${RELEASE_NAME}-env-local --from-env-file=.env --dry-run=client -o yaml | $(KUBECTL_EXEC) apply -f -
#
#.PHONY: env.secret.create
#env.secret.create:
#	$(KUBECTL_EXEC) create secret generic ${RELEASE_NAME}-env-secret-local --from-literal=username=devuser
#
#.PHONY: deploy_k8s
#deploy_k8s: env.create env.secret.create ## Деплой kubernetes
#	helm upgrade --install \
#		${RELEASE_NAME} \
#		${HELM_DIR} \
#		-f ${HELM_DIR}/values.local.yaml \
#		--namespace=default \
#		--set global.image=${IMAGE}:${TAG} \
#		--debug
#
#.PHONY: delete_k8s
#delete_k8s: ## Удаление деплоя kubernetes
#	helm delete ${RELEASE_NAME}
#
#mkube.up:
#	minikube start --kubernetes-version v${K8S_VERSION} --mount-string="${GOPATH}/src/${PROJECT}:/go/src/${PROJECT}" --mount
#
#mkube.del:
#	minikube delete
