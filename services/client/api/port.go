package api

import (
	"client/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type PortHandler struct {
	portService services.PortServiceInterface
}

func NewPortHandler(portService services.PortServiceInterface) *PortHandler {
	service := &PortHandler{
		portService: portService,
	}
	return service
}

func (p *PortHandler) getPort(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	key := vars["id"]
	port, err := p.portService.GetPort(ctx, key)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if port == nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(port)
}

func (p *PortHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/ports/{id}", p.getPort)
}