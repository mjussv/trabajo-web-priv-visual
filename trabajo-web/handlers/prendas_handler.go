package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"trabajo-web/logic"
)

type PrendasHandler struct {
	prendasService *logic.PrendasService
}

func NewPrendasHandler(prendasService *logic.PrendasService) *PrendasHandler {
	return &PrendasHandler{
		prendasService: prendasService,
	}
}

// POST /prendas
func (h *PrendasHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req logic.CreatePrendasParams

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Formato JSON inválido", http.StatusBadRequest)
		return
	}

	prendas, err := h.prendasService.CreatePrendas(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(prendas)
}

// GET /prendas
func (h *PrendasHandler) List(w http.ResponseWriter, r *http.Request) {
	prendas, err := h.prendasService.ListPrendas(r.Context())
	if err != nil {
		http.Error(w, "Error interno al obtener prendas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prendas)
}

// GET /prendas/{id}
func (h *PrendasHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	prendas, err := h.prendasService.GetPrendasByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Prendas no encontradas", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prendas)
}

// DELETE /prendas/{id}
func (h *PrendasHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.prendasService.DeletePrendas(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
