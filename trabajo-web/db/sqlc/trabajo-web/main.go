package trabajoweb
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	// Driver pgx v5 para PostgreSQL
	_ "github.com/jackc/pgx/v5/stdlib"

	"trabajo-web/db/sqlc"
	"trabajo-web/handlers"
	"trabajo-web/logic"
)

func main() {
	// 1. Obtener la cadena de conexión a PostgreSQL
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/ootd_db?sslmode=disable"
	}

	// 2. Conectar a PostgreSQL usando el driver pgx
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("Error configurando la base de datos: %v\n", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v\n", err)
	}
	log.Println("Conexión exitosa a la base de datos PostgreSQL")

	// 3. Inicializar sqlc Queries
	queries := sqlc.New(db)

	// 4. Inicializar Capa de Lógica (Services)
	usuariosService := logic.NewUsuariosService(queries)
	prendasService := logic.NewPrendasService(queries)
	looksService := logic.NewLooksService(queries)
	meGustaLooksService := logic.NewMeGustaLooksService(queries)

	// 5. Inicializar Capa de Presentación (Handlers)
	usuariosHandler := handlers.NewUsuariosHandler(usuariosService)
	prendasHandler := handlers.NewPrendasHandler(prendasService)
	looksHandler := handlers.NewLooksHandler(looksService)
	meGustaLooksHandler := handlers.NewMeGustaLooksHandler(meGustaLooksService)

	// 6. Obtener el router con todas las rutas configuradas
	router := handlers.SetupRoutes(usuariosHandler, prendasHandler, looksHandler, meGustaLooksHandler)

	// 7. Servidor de archivos estáticos (Opcional, si necesitás frontend/estáticos)
	// Nota: Si usás el router de Go 1.22+, pods registrar la carpeta estática dentro de handlers/routes.go
	staticDir := "./static"

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Servidor escuchando en http://localhost:%s\n", port)
	fmt.Printf("Sirviendo archivos desde: %s\n", staticDir)

	// 8. Iniciar el servidor pasando nuestro 'router'
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Error al iniciar el servidor: %s\n", err)
	}
}