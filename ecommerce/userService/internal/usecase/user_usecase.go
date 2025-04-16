package usecase

import (
	"userService/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) domain.UserUsecase {
	return &userUsecase{r}
}

func (uc *userUsecase) Register(username, password string) (*domain.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &domain.User{
		Username: username,
		Password: string(hash),
	}
	err = uc.repo.Create(u)
	return u, err
}

func (uc *userUsecase) Authenticate(username, password string) (*domain.User, error) {
	u, err := uc.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *userUsecase) GetProfile(id int) (*domain.User, error) {
	return uc.repo.GetByID(id)
}
