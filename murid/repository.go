package murid

import (
	"gorm.io/gorm"
)

type MuridRepository struct {
	db *gorm.DB
}

func NewMuridRepository(db *gorm.DB) *MuridRepository {
	return &MuridRepository{db}
}

func (r *MuridRepository) CreateMurid(murid *Murid) (err error) {
	err = r.db.Create(murid).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *MuridRepository) GetAllMurid() ([]Murid, error) {
	var murids []Murid
	err := r.db.Find(&murids).Error
	if err != nil {
		return murids, err
	}
	return murids, err
}

func (r *MuridRepository) UpdateMurid(murid *Murid) error {
	err := r.db.Save(murid).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MuridRepository) DeleteMurid(id uint) error {
	err := r.db.Delete(&Murid{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
