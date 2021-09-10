package service

import (
	"avito-intern/internal/models"
	"avito-intern/internal/store/postgres"
)

type UserMoneyService interface {
	GetUserMoneyAmount(uuid string) (*models.UserMoney, error)
	// CreateUser(userName string) (*models.User, error)
	// UpdateUser(user *models.User) (*models.User, error)
	// DeleteUser(uuid string) (bool, error)
}

type UserMoneyServiceCases struct {
	MoneyAmountRepo UserMoneyService
}

func NewUserServiceCases() *UserMoneyServiceCases {
	return &UserMoneyServiceCases{
		MoneyAmountRepo: postgres.NewMoneyServiceRepo(),
	}
}

func (u UserMoneyServiceCases) GetUserMoneyAmount(uuid string) (*models.UserMoney, error) {
	return u.MoneyAmountRepo.GetUserMoneyAmount(uuid)
}

// func (u UserMoneyServiceCases) CreateUser(userName string) (*models.User, error) {
// 	return u.MoneyAmountRepo.CreateUser(userName)
// }

// func (u UserMoneyServiceCases) UpdateUser(user *models.User) (*models.User, error) {
// 	return u.MoneyAmountRepo.UpdateUser(user)
// }

// func (u UserMoneyServiceCases) DeleteUser(uuid string) (bool, error) {
// 	return u.MoneyAmountRepo.DeleteUser(uuid)
// }
