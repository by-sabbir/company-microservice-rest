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
	private := h.Router.PathPrefix("/api/v1/private/company").Subrouter()
	public := h.Router.PathPrefix("/api/v1/public/company").Subrouter()

	public.HandleFunc("/{id}", h.GetCompany).Methods("GET")

	private.HandleFunc("/create", h.PostCompany).Methods("POST")
	private.HandleFunc("/delete/{id}", h.DeleteCompany).Methods("DELETE")
	private.HandleFunc("/patch/{id}", h.PartialUpdateCompany).Methods("PATCH")

}

// go test - skip
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Fatalf("error serving http%+v\n", err)
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
