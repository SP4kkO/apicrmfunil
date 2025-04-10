package historicoetapa

import (
	"errors"

	"gorm.io/gorm"
)

// Repository define as operações básicas para manipulação de históricos de etapa.
type Repository interface {
	Adicionar(h HistoricoEtapa) (HistoricoEtapa, error)
	Listar() ([]HistoricoEtapa, error)
	ObterPorID(id int) (HistoricoEtapa, error)
	Atualizar(id int, h HistoricoEtapa) (HistoricoEtapa, error)
	Deletar(id int) error
	ListarPorNegociacao(negociacaoId int) ([]HistoricoEtapa, error)
}

type repository struct {
	db *gorm.DB
}

// NovoRepositorio cria um novo repositório para Histórico de Etapas.
func NovoRepositorio(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Adicionar(h HistoricoEtapa) (HistoricoEtapa, error) {
	err := r.db.Create(&h).Error
	return h, err
}

func (r *repository) Listar() ([]HistoricoEtapa, error) {
	var historicos []HistoricoEtapa
	err := r.db.Find(&historicos).Error
	return historicos, err
}

func (r *repository) ObterPorID(id int) (HistoricoEtapa, error) {
	var h HistoricoEtapa
	err := r.db.First(&h, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return h, errors.New("histórico não encontrado")
		}
		return h, err
	}
	return h, nil
}

func (r *repository) Atualizar(id int, h HistoricoEtapa) (HistoricoEtapa, error) {
	var historico HistoricoEtapa
	err := r.db.First(&historico, id).Error
	if err != nil {
		return historico, errors.New("histórico não encontrado")
	}
	// Garanta que o ID não seja modificado
	h.ID = historico.ID
	err = r.db.Model(&historico).Updates(h).Error
	return h, err
}

func (r *repository) Deletar(id int) error {
	return r.db.Delete(&HistoricoEtapa{}, id).Error
}

func (r *repository) ListarPorNegociacao(negociacaoId int) ([]HistoricoEtapa, error) {
	var historicos []HistoricoEtapa
	err := r.db.Where("negociacao_id = ?", negociacaoId).Find(&historicos).Error
	return historicos, err
}
