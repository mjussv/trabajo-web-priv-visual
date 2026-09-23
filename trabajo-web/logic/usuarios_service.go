package logic

import (
	"context"
	"errors"
	"strings"

	"trabajo-web/db/sqlc"
)

type UsuariosService struct {
	queries *sqlc.Queries
}

func NewUsuariosService(queries *sqlc.Queries) *UsuariosService {
	return &UsuariosService{
		queries: queries,
	}
}

type CreateUsuariosParams struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

func (s *UsuariosService) CreateUsuarios(ctx context.Context, params CreateUsuariosParams) (sqlc.Usuario, error) {
	if strings.TrimSpace(params.Nombre) == "" {
		return sqlc.Usuario{}, errors.New("el nombre no puede estar vacío")
	}
	if strings.TrimSpace(params.Email) == "" {
		return sqlc.Usuario{}, errors.New("el email no puede estar vacío")
	}

	return s.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Nombre: params.Nombre,
		Email:  params.Email,
	})
}

func (s *UsuariosService) ListUsuarios(ctx context.Context) ([]sqlc.Usuario, error) {
	return s.queries.ListUsers(ctx)
}

func (s *UsuariosService) GetUsuariosByID(ctx context.Context, id int64) (sqlc.Usuario, error) {
	if id <= 0 {
		return sqlc.Usuario{}, errors.New("el ID de usuario debe ser mayor a 0")
	}
	return s.queries.GetUser(ctx, id)
}

type UpdateUsuariosParams struct {
	ID     int64  `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

func (s *UsuariosService) UpdateUsuarios(ctx context.Context, params UpdateUsuariosParams) (sqlc.Usuario, error) {
	if params.ID <= 0 {
		return sqlc.Usuario{}, errors.New("el ID de usuario debe ser mayor a 0")
	}
	if strings.TrimSpace(params.Nombre) == "" {
		return sqlc.Usuario{}, errors.New("el nombre no puede estar vacío")
	}

	return s.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:     params.ID,
		Nombre: params.Nombre,
		Email:  params.Email,
	})
}

func (s *UsuariosService) DeleteUsuarios(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("el ID de usuario debe ser mayor a 0")
	}
	return s.queries.DeleteUser(ctx, id)
}
