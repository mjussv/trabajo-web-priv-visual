package logic

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"trabajo-web/db/sqlc"
)

type LooksService struct {
	queries *sqlc.Queries
}

func NewLooksService(queries *sqlc.Queries) *LooksService {
	return &LooksService{queries: queries}
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
	if strings.TrimSpace(params.Ocasion) == "" {
		return sqlc.Look{}, errors.New("la ocasión no puede estar vacía")
	}
	if strings.TrimSpace(params.Temporada) == "" {
		return sqlc.Look{}, errors.New("la temporada no puede estar vacía")
	}
	if strings.TrimSpace(params.UrlImagen) == "" {
		return sqlc.Look{}, errors.New("la URL de la imagen no puede estar vacía")
	}

	return s.queries.CrearLook(ctx, sqlc.CrearLookParams{
		UsuarioID:   params.UsuarioID,
		Nombre:      params.Nombre,
		Descripcion: sql.NullString{String: params.Descripcion, Valid: params.Descripcion != ""},
		Ocasion:     params.Ocasion,
		Temporada:   params.Temporada,
		PrendaEstrellaID: sql.NullInt32{
			Int32: params.PrendaEstrellaID,
			Valid: params.PrendaEstrellaID != 0,
		},
		UrlImagen: params.UrlImagen,
		Estilo:    sql.NullString{String: params.Estilo, Valid: params.Estilo != ""},
	})
}

func (s *LooksService) ListLooks(ctx context.Context) ([]sqlc.ListarLooksConDetalleRow, error) {
	return s.queries.ListarLooksConDetalle(ctx)
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
