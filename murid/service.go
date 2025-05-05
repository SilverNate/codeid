package murid

import "context"

type MuridService struct {
	repo IMuridRepo
}

func NewMuridService(repo IMuridRepo) *MuridService {
	return &MuridService{repo: repo}
}

func (s *MuridService) CreateMurid(ctx context.Context, request Murid) (err error) {
	return s.repo.CreateMurid(&request)
}

func (s *MuridService) GetMurid(ctx context.Context) (result []Murid, err error) {
	return s.repo.GetAllMurid()
}
