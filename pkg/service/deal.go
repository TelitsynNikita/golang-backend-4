package service

import (
	todo "github.com/TelitsynNikita"
	"github.com/TelitsynNikita/pkg/repository"
)

type DealService struct {
	repo repository.Deals
}

func newTodoDealService(repo repository.Deals) *DealService {
	return &DealService{repo: repo}
}

func (s *DealService) Create(userId int, deal todo.Deal) (int, error) {
	return s.repo.Create(userId, deal)
}

func (s *DealService) GetAllNew() ([]todo.AllNewDeals, error) {
	return s.repo.GetAllNew()
}

func (s *DealService) GetOneDealById(id int) (todo.OneDeal, error) {
	return s.repo.GetOneDealById(id)
}

func (s *DealService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *DealService) UpdateStatus(status string, id int) error {
	return s.repo.UpdateStatus(status, id)
}

func (s *DealService) UpdateDealBookkeeperId(userId int, dealId int) error {
	return s.repo.UpdateDealBookkeeperId(userId, dealId)
}

func (s *DealService) GetAllOwnDeals(id int, role string, status string) ([]todo.AllNewDeals, error) {
	return s.repo.GetAllOwnDeals(id, role, status)
}
