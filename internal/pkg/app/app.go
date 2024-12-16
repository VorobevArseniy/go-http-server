package app

import (
	"fmt"

	"go-http-server/internal/app/endpoint"
	middleware "go-http-server/internal/app/mw"
	"go-http-server/internal/app/service"
	"log"
	"net/http"
)

const (
	portNum string = ":8080"
)

type App struct {
	e   *endpoint.Endpoint
	s   *service.Service
	r   *http.ServeMux
	srv *http.Server
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.r = http.NewServeMux()

	stack := middleware.CreateStack(
		middleware.Logger,
	)

	a.srv = &http.Server{
		Addr:    portNum,
		Handler: stack(a.r),
	}

	a.r.HandleFunc("GET /user/{name}", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Printf("server runing on port: %v ~ \n", portNum)

	err := a.srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
