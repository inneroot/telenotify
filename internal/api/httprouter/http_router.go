package httpRouter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	notify_service "github.com/inneroot/telenotify/internal/service"
)

type Router struct {
	routes map[string]http.HandlerFunc
	ns     *notify_service.NotifyService
}

func New(ns *notify_service.NotifyService) *Router {
	router := Router{
		routes: make(map[string]http.HandlerFunc),
		ns:     ns,
	}
	router.AddRoute("/", router.handleNotify)
	return &router
}

func (r *Router) AddRoute(path string, handler http.HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Check if the path exists in the routes map
	if handler, exists := r.routes[req.URL.Path]; exists {
		handler(w, req)
	} else {
		// Handle 404 Not Found
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}
}

func (r *Router) handleNotify(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		slog.Error("http notify: error reading request body", slog.String("error", err.Error()))
		http.Error(w, "error reading request body", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	notification := NotifyPostRequestBody{}
	if err := json.Unmarshal(body, &notification); err != nil {
		slog.Error("http notify: unmarshal request body", slog.String("error", err.Error()))
		http.Error(w, "error reading request body", http.StatusBadRequest)
		return
	}

	// TODO: sign verification

	slog.Info("notify", slog.String("message", notification.Message))
	if err := r.ns.Notify(context.Background(), notification.Message); err != nil {
		slog.Error("http notify: unmarshal request body", slog.String("error", err.Error()))
		http.Error(w, "error reading request body", http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "message have bin sent")
}
