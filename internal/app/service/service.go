package service

import (
	"math/rand"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) RandomCoupon() int64 {
	// some buisness logic here

	money := rand.Intn(1000) // for example we get wallet balance from another service

	return int64(money)
}
