package logic

import (
	"context"

	"trabajo-web/db/sqlc"
)

type MeGustaLooksService struct {
	queries *sqlc.Queries
}

func NewMeGustaLooksService(queries *sqlc.Queries) *MeGustaLooksService {
	return &MeGustaLooksService{queries: queries}
}

type CreateMeGustaParams struct {
	UsuarioID int32 `json:"usuario_id"`
	LookID    int32 `json:"look_id"`
}

func (s *MeGustaLooksService) AddMeGusta(ctx context.Context, params CreateMeGustaParams) error {
	return s.queries.DarMeGusta(ctx, sqlc.DarMeGustaParams{
		UsuarioID: params.UsuarioID,
		LookID:    params.LookID,
	})
}

func (s *MeGustaLooksService) DeleteMeGusta(ctx context.Context, params CreateMeGustaParams) error {
	return s.queries.QuitarMeGusta(ctx, sqlc.QuitarMeGustaParams{
		UsuarioID: params.UsuarioID,
		LookID:    params.LookID,
	})
}
