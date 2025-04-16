package usecase

import "orderService/internal/domain"

type orderUsecase struct {
	repo domain.OrderRepository
}

func NewOrderUsecase(r domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{r}
}

func (uc *orderUsecase) Create(o *domain.Order) error {
	o.Status = "pending"
	return uc.repo.Create(o)
}

func (uc *orderUsecase) GetByID(id int) (*domain.Order, error) {
	return uc.repo.GetByID(id)
}

func (uc *orderUsecase) UpdateStatus(id int, status string) error {
	return uc.repo.UpdateStatus(id, status)
}

func (uc *orderUsecase) ListByUser(userID int) ([]domain.Order, error) {
	return uc.repo.ListByUser(userID)
}
