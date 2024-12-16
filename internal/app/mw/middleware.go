package middleware

import (
	"context"
	"log"
	"net/http"
	"time"
)

const (
	UserRole  = "middleware.Logging.Role"
	roleAdmin = "admin"
)

type Middleware struct{ UserRole string }

func New() *Middleware {
	return &Middleware{UserRole: UserRole}
}

type mw func(http.Handler) http.Handler

type wrapperWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrapperWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func (m *Middleware) CreateStack(xs ...mw) mw {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

func (m *Middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrapperWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		role := r.Header.Get("role")

		log.Println(wrapped.statusCode, r.Method, r.URL.Path, role, time.Since(start))

		ctx := context.WithValue(r.Context(), UserRole, role)
		req := r.WithContext(ctx)

		next.ServeHTTP(wrapped, req)
	})
}

// func RoleCheck(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		if val := r.Header.Get("User-Role"); strings.Contains(val, roleAdmin) {
// 			log.Println("red button user detected")
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }
