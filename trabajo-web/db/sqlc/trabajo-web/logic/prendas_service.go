package logic

import (
	"context"
	"errors"
	"strings"

	"trabajo-web/db/sqlc"
)

type PrendasService struct {
	queries *sqlc.Queries
}

func NewPrendasService(queries *sqlc.Queries) *PrendasService {
	return &PrendasService{
		queries: queries,
	}
}

type CreatePrendasParams struct {
	Nombre    string `json:"nombre"`
	Categoria string `json:"categoria"`
}

func (s *PrendasService) CreatePrendas(ctx context.Context, params CreatePrendasParams) (sqlc.Prenda, error) {
	if strings.TrimSpace(params.Nombre) == "" {
		return sqlc.Prenda{}, errors.New("el nombre de la prenda no puede estar vacío")
	}

	return s.queries.CreatePrenda(ctx, sqlc.CreatePrendaParams{
		Nombre:    params.Nombre,
		Categoria: params.Categoria,
	})
}

func (s *PrendasService) ListPrendas(ctx context.Context) ([]sqlc.Prenda, error) {
	return s.queries.ListPrendas(ctx)
}

func (s *PrendasService) GetPrendasByID(ctx context.Context, id int64) (sqlc.Prenda, error) {
	if id <= 0 {
		return sqlc.Prenda{}, errors.New("ID inválido")
	}
	return s.queries.GetPrenda(ctx, id)
}

func (s *PrendasService) DeletePrendas(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("ID inválido")
	}
	return s.queries.DeletePrenda(ctx, id)
}
