package usecase

import "inventoryService/internal/domain"

// 8. Usecase ссылается на product_repo.go
type productUsecase struct {
	repo domain.ProductRepository
}

// 8.1 Передавая ему "правила работы" с конкретными методами
func NewProductUsecase(r domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{r}
}

func (uc *productUsecase) Create(p *domain.Product) error {
	return uc.repo.Create(p)
}

func (uc *productUsecase) GetByID(id int) (*domain.Product, error) {
	return uc.repo.GetByID(id)
}

func (uc *productUsecase) Update(p *domain.Product) error {
	return uc.repo.Update(p)
}

func (uc *productUsecase) Delete(id int) error {
	return uc.repo.Delete(id)
}

func (uc *productUsecase) List() ([]domain.Product, error) {
	return uc.repo.List()
}
