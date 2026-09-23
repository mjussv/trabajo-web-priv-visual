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

	return s.queries.CrearUsuario(ctx, sqlc.CrearUsuarioParams{
		Nombre: params.Nombre,
		Email:  params.Email,
	})
}

func (s *UsuariosService) ListUsuarios(ctx context.Context) ([]sqlc.Usuario, error) {
	return s.queries.ListarUsuarios(ctx)
}

func (s *UsuariosService) GetUsuariosByID(ctx context.Context, id int32) (sqlc.Usuario, error) {
	if id <= 0 {
		return sqlc.Usuario{}, errors.New("el ID de usuario debe ser mayor a 0")
	}
	return s.queries.ObtenerUsuario(ctx, id)
}
