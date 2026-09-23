package handlers

package handlers

import (
	"net/http"
)

func SetupRoutes(
	usuariosHandler *UsuariosHandler,
	prendasHandler *PrendasHandler,
	looksHandler *LooksHandler,
	meGustaLooksHandler *MeGustaLooksHandler,
) http.Handler {
	mux := http.NewServeMux()

	// Servir archivos estáticos
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	// Rutas de Usuarios
	mux.HandleFunc("POST /usuarios", usuariosHandler.Create)
	mux.HandleFunc("GET /usuarios", usuariosHandler.List)
	mux.HandleFunc("GET /usuarios/{id}", usuariosHandler.GetByID)
	mux.HandleFunc("PUT /usuarios/{id}", usuariosHandler.Update)
	mux.HandleFunc("DELETE /usuarios/{id}", usuariosHandler.Delete)

	// Rutas de Prendas
	mux.HandleFunc("POST /prendas", prendasHandler.Create)
	mux.HandleFunc("GET /prendas", prendasHandler.List)
	mux.HandleFunc("GET /prendas/{id}", prendasHandler.GetByID)
	mux.HandleFunc("DELETE /prendas/{id}", prendasHandler.Delete)

	// Rutas de Looks
	mux.HandleFunc("POST /looks", looksHandler.Create)
	mux.HandleFunc("GET /looks", looksHandler.List)
	mux.HandleFunc("GET /looks/{id}", looksHandler.GetByID)
	mux.HandleFunc("DELETE /looks/{id}", looksHandler.Delete)

	// Rutas de Me Gusta
	mux.HandleFunc("POST /me-gusta", meGustaLooksHandler.AddMeGusta)
	mux.HandleFunc("DELETE /me-gusta", meGustaLooksHandler.DeleteMeGusta)

	return mux
}