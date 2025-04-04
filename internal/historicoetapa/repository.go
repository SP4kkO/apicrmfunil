package historicoetapa

import "gorm.io/gorm"

type Repository interface {
	Adicionar(h HistoricoEtapa) (HistoricoEtapa, error)
	Listar() ([]HistoricoEtapa, error)
	ObterPorID(id int) (*HistoricoEtapa, error)
	Atualizar(id int, updated HistoricoEtapa) (HistoricoEtapa, error)
	Deletar(id int) error
	ListarPorNegociacao(negociacaoID int) ([]HistoricoEtapa, error)
}

type repository struct {
	db *gorm.DB
}

func NovoRepositorio(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Adicionar(h HistoricoEtapa) (HistoricoEtapa, error) {
	err := r.db.Create(&h).Error
	return h, err
}

func (r *repository) Listar() ([]HistoricoEtapa, error) {
	var historicos []HistoricoEtapa
	err := r.db.Order("data_alteracao DESC").Find(&historicos).Error
	return historicos, err
}

func (r *repository) ObterPorID(id int) (*HistoricoEtapa, error) {
	var historico HistoricoEtapa
	err := r.db.First(&historico, id).Error
	if err != nil {
		return nil, err
	}
	return &historico, nil
}

func (r *repository) Atualizar(id int, updated HistoricoEtapa) (HistoricoEtapa, error) {
	var historico HistoricoEtapa
	err := r.db.First(&historico, id).Error
	if err != nil {
		return HistoricoEtapa{}, err
	}

	updated.ID = historico.ID
	err = r.db.Model(&historico).Updates(updated).Error
	return updated, err
}

func (r *repository) Deletar(id int) error {
	return r.db.Delete(&HistoricoEtapa{}, id).Error
}

func (r *repository) ListarPorNegociacao(negociacaoID int) ([]HistoricoEtapa, error) {
	var historicos []HistoricoEtapa
	err := r.db.Where("negociacao_id = ?", negociacaoID).
		Order("data_alteracao DESC").
		Find(&historicos).Error
	return historicos, err
}
