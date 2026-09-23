package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"trabajo-web/logic"
)

type UsuariosHandler struct {
	usuariosService *logic.UsuariosService
}

func NewUsuariosHandler(usuariosService *logic.UsuariosService) *UsuariosHandler {
	return &UsuariosHandler{
		usuariosService: usuariosService,
	}
}

// POST /usuarios
func (h *UsuariosHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req logic.CreateUsuariosParams

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Formato JSON inválido", http.StatusBadRequest)
		return
	}

	usuario, err := h.usuariosService.CreateUsuarios(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

// GET /usuarios
func (h *UsuariosHandler) List(w http.ResponseWriter, r *http.Request) {
	usuarios, err := h.usuariosService.ListUsuarios(r.Context())
	if err != nil {
		http.Error(w, "Error interno al obtener usuarios", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// GET /usuarios/{id}
func (h *UsuariosHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	usuario, err := h.usuariosService.GetUsuariosByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

// PUT /usuarios/{id}
func (h *UsuariosHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var req logic.UpdateUsuariosParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Formato JSON inválido", http.StatusBadRequest)
		return
	}

	req.ID = id

	usuario, err := h.usuariosService.UpdateUsuarios(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

// DELETE /usuarios/{id}
func (h *UsuariosHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.usuariosService.DeleteUsuarios(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
