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
	UsuarioID        int32  `json:"usuario_id"`
	Nombre           string `json:"nombre"`
	Descripcion      string `json:"descripcion"`
	Ocasion          string `json:"ocasion"`
	Temporada        string `json:"temporada"`
	PrendaEstrellaID int32  `json:"prenda_estrella_id"`
	UrlImagen        string `json:"url_imagen"`
	Estilo           string `json:"estilo"`
}

func (s *LooksService) CreateLooks(ctx context.Context, params CreateLooksParams) (sqlc.Look, error) {
	if strings.TrimSpace(params.Nombre) == "" {
		return sqlc.Look{}, errors.New("el nombre del look no puede estar vacío")
	}

	return s.queries.CrearLook(ctx, sqlc.CrearLookParams{
		UsuarioID:        params.UsuarioID,
		Nombre:           params.Nombre,
		Descripcion:      params.Descripcion,
		Ocasion:          params.Ocasion,
		Temporada:        params.Temporada,
		PrendaEstrellaID: params.PrendaEstrellaID,
		UrlImagen:        params.UrlImagen,
		Estilo:           params.Estilo,
	})
}

func (s *LooksService) GetLooksByID(ctx context.Context, id int32) (sqlc.ObtenerLookConDetalleRow, error) {
	if id <= 0 {
		return sqlc.ObtenerLookConDetalleRow{}, errors.New("ID inválido")
	}
	return s.queries.ObtenerLookConDetalle(ctx, id)
}

func (s *LooksService) DeleteLooks(ctx context.Context, id int32) error {
	if id <= 0 {
		return errors.New("ID inválido")
	}
	return s.queries.EliminarLook(ctx, id)
}
