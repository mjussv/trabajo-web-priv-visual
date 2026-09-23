package logic

import (
	"context"
	"errors"
	"strings"

	"trabajo-web/db/sqlc"
)

type LooksService struct {
	queries *sqlc.Queries
}

func NewLooksService(queries *sqlc.Queries) *LooksService {
	return &LooksService{
		queries: queries,
	}
}

type CreateLooksParams struct {
	Titulo    string `json:"titulo"`
	UsuarioID int64  `json:"usuario_id"`
}

func (s *LooksService) CreateLooks(ctx context.Context, params CreateLooksParams) (sqlc.Look, error) {
	if strings.TrimSpace(params.Titulo) == "" {
		return sqlc.Look{}, errors.New("el título del look no puede estar vacío")
	}
	if params.UsuarioID <= 0 {
		return sqlc.Look{}, errors.New("el ID de usuario es obligatorio")
	}

	return s.queries.CreateLook(ctx, sqlc.CreateLookParams{
		Titulo:    params.Titulo,
		UsuarioID: params.UsuarioID,
	})
}

func (s *LooksService) ListLooks(ctx context.Context) ([]sqlc.Look, error) {
	return s.queries.ListLooks(ctx)
}

func (s *LooksService) GetLooksByID(ctx context.Context, id int64) (sqlc.Look, error) {
	if id <= 0 {
		return sqlc.Look{}, errors.New("ID inválido")
	}
	return s.queries.GetLook(ctx, id)
}

func (s *LooksService) DeleteLooks(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("ID inválido")
	}
	return s.queries.DeleteLook(ctx, id)
}
