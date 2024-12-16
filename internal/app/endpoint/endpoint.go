package endpoint

import (
	"fmt"
	middleware "go-http-server/internal/app/mw"
	"log"
	"net/http"
	"strings"
)

type Service interface {
	RandomCoupon() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(w http.ResponseWriter, r *http.Request) {
	m := middleware.New()
	n := r.PathValue("name")
	c := e.s.RandomCoupon()

	role, ok := r.Context().Value(m.UserRole).(string)
	if !ok {
		log.Println("Invalid user role")
	}

	var s string

	if strings.Contains(strings.ToLower(role), "admin") {
		s = fmt.Sprintf("Hello, %v. You have admin role; *show admin panel* \n", n)
	} else {
		s = fmt.Sprintf("Hello, %v. You have received a new coupon! Coupon amount: %v \n", n, c)
	}

	res := []byte(s)
	if _, err := w.Write(res); err != nil {
		log.Printf("Failed to write response: %v", err)
	}

}
