package murid

import "context"

type IMuridService interface {
	GetMurid(ctx context.Context) (resp []Murid, err error)
	CreateMurid(ctx context.Context, request Murid) (err error)
	//UpdateMurid(ctx context.Context, request Murid) (err error)
	//DeleteMurid(ctx context.Context) (err error)
}

type IMuridRepo interface {
	CreateMurid(murid *Murid) (err error)
	GetAllMurid() ([]Murid, error)
	UpdateMurid(murid *Murid) error
	DeleteMurid(id uint) error
}
