package logic

import (
	"context"
	"errors"

	"trabajo-web/db/sqlc"
)

type MeGustaLooksService struct {
	queries *sqlc.Queries
}

func NewMeGustaLooksService(queries *sqlc.Queries) *MeGustaLooksService {
	return &MeGustaLooksService{
		queries: queries,
	}
}

type MeGustaParams struct {
	UsuarioID int32 `json:"usuario_id"`
	LookID    int32 `json:"look_id"`
}

func (s *MeGustaLooksService) AddMeGusta(ctx context.Context, params MeGustaParams) error {
	if params.UsuarioID <= 0 || params.LookID <= 0 {
		return errors.New("los ID de usuario y look deben ser válidos")
	}

	return s.queries.DarMeGusta(ctx, sqlc.DarMeGustaParams{
		UsuarioID: params.UsuarioID,
		LookID:    params.LookID,
	})
}

func (s *MeGustaLooksService) DeleteMeGusta(ctx context.Context, params MeGustaParams) error {
	if params.UsuarioID <= 0 || params.LookID <= 0 {
		return errors.New("los ID de usuario y look deben ser válidos")
	}

	return s.queries.QuitarMeGusta(ctx, sqlc.QuitarMeGustaParams{
		UsuarioID: params.UsuarioID,
		LookID:    params.LookID,
	})
}
