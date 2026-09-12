package db_test

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	sqlc "trabajo-web/db/sqlc" // Si tu go.mod no se llama trabajo-web, cambialo acá
)

const (
	connStr = "postgres://postgres:mysecretpassword@localhost:5432/outfits_db?sslmode=disable"
)

func TestQueries_CRUD(t *testing.T) {
	ctx := context.Background()

	// 1. Conexión a la base de datos PostgreSQL
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		t.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatalf("La base de datos no está respondiendo: %v", err)
	}

	// 2. Instanciar cliente generado por sqlc
	queries := sqlc.New(db)

	var createdUserID int32

	// Step 1: CrearUsuario
	t.Run("CrearUsuario", func(t *testing.T) {
		createdUser, err := queries.CrearUsuario(ctx, sqlc.CrearUsuarioParams{
			Nombre: "John Doe",
			Email:  "john.doe@example.com",
		})
		if err != nil {
			t.Fatalf("failed to create user: %v", err)
		}

		if createdUser.ID == 0 {
			t.Errorf("Se esperaba un ID válido pero se obtuvo 0")
		}

		createdUserID = createdUser.ID
		t.Logf("Created user: %+v", createdUser)
	})

	if createdUserID == 0 {
		t.Fatal("Imposible continuar las pruebas CRUD sin un usuario creado")
	}

	// Step 2: ObtenerUsuario (Read One)
	t.Run("ObtenerUsuario", func(t *testing.T) {
		user, err := queries.ObtenerUsuario(ctx, createdUserID)
		if err != nil {
			t.Fatalf("failed to get user: %v", err)
		}

		if user.ID != createdUserID {
			t.Errorf("Se esperaba ID %d, pero se obtuvo %d", createdUserID, user.ID)
		}
		t.Logf("Retrieved user: %+v", user)
	})

	// Step 3: ListarUsuarios (Read Many)
	t.Run("ListarUsuarios", func(t *testing.T) {
		users, err := queries.ListarUsuarios(ctx)
		if err != nil {
			t.Fatalf("failed to list users: %v", err)
		}

		if len(users) == 0 {
			t.Errorf("Se esperaba al menos 1 usuario en la lista")
		}
		t.Logf("All users: %+v", users)
	})
}
