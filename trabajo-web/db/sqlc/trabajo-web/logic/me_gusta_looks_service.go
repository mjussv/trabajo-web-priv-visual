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

type CreateMeGustaParams struct {
	UsuarioID int64 `json:"usuario_id"`
	LookID    int64 `json:"look_id"`
}

func (s *MeGustaLooksService) AddMeGusta(ctx context.Context, params CreateMeGustaParams) (sqlc.MeGustaLook, error) {
	if params.UsuarioID <= 0 || params.LookID <= 0 {
		return sqlc.MeGustaLook{}, errors.New("los ID de usuario y look deben ser válidos")
	}

	return s.queries.CreateMeGustaLook(ctx, sqlc.CreateMeGustaLookParams{
		UsuarioID: params.UsuarioID,
		LookID:    params.LookID,
	})
}

func (s *MeGustaLooksService) DeleteMeGusta(ctx context.Context, params CreateMeGustaParams) error {
	if params.UsuarioID <= 0 || params.LookID <= 0 {
		return errors.New("los ID de usuario y look deben ser válidos")
	}

	return s.queries.DeleteMeGustaLook(ctx, sqlc.DeleteMeGustaLookParams{
		UsuarioID: params.UsuarioID,
		LookID:    params.LookID,
	})
}
