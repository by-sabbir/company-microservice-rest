package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service CompanyRestService
	Server  *http.Server
}

func NewHandler(service CompanyRestService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Router.Use(LogMiddlewire)
	h.Router.Use(JSONMiddlewire)

	h.Server = &http.Server{
		Addr:         "0.0.0.0:8888",
		Handler:      h.Router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return h
}

func (h *Handler) mapRoutes() {
	private := h.Router.PathPrefix("/api/v1/private").Subrouter()
	public := h.Router.PathPrefix("/api/v1/public").Subrouter()

	public.HandleFunc("/company/{id}", h.GetCompany)

	private.HandleFunc("/create-company", h.PostCompany)

}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("Shut down gracefully...")
	return nil
}