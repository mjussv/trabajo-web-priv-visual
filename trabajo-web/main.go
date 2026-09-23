package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"trabajo-web/db/sqlc"
	"trabajo-web/handlers"
	"trabajo-web/logic"
)

func main() {
	// 1. Obtener URL de conexión
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/outfits_db?sslmode=disable"
	}

	// 2. Conectar a la base de datos
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("Error configurando la base de datos: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	// 3. Capa DB (sqlc)
	queries := sqlc.New(db)

	// 4. Capa Lógica (Services)
	usuariosSvc := logic.NewUsuariosService(queries)
	prendasSvc := logic.NewPrendasService(queries)
	looksSvc := logic.NewLooksService(queries)
	meGustaSvc := logic.NewMeGustaLooksService(queries)

	// 5. Capa HTTP (Handlers)
	usuariosHdl := handlers.NewUsuariosHandler(usuariosSvc)
	prendasHdl := handlers.NewPrendasHandler(prendasSvc)
	looksHdl := handlers.NewLooksHandler(looksSvc)
	meGustaHdl := handlers.NewMeGustaLooksHandler(meGustaSvc)

	// 6. Configurar Rutas
	router := handlers.SetupRoutes(usuariosHdl, prendasHdl, looksHdl, meGustaHdl)

	// 7. Iniciar Servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Servidor corriendo en http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}
