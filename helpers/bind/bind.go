package bind

import "strconv"

type SequentialGenerator struct {
	last int
}

func NewSequentialGenerator(start int) *SequentialGenerator {
	return &SequentialGenerator{start}
}

func (s *SequentialGenerator) Next() (res int) {
	res = s.last
	s.last++
	return
}

func (s *SequentialGenerator) NextBind() (res string) {
	res = "$" + strconv.FormatInt(int64(s.last), 10)
	s.last++
	return
}
