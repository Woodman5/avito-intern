package service

import (
	"avito-intern/internal/models"
	"avito-intern/internal/store/postgres"
	"avito-intern/internal/utils"
)

type UserMoneyService interface {
	GetUserMoneyAmount(uuid string) (*models.UserMoneyFloat, error)
	CreateTransaction(request *models.TransactionRequest) (*models.UserMoneyFloat, error)
	CreateUser(user *models.UserMoneyFloat) (*models.UserMoneyFloat, error)
	FundsTransfer(request *models.TransferRequest) (*models.TransferAnswer, error)
}

type UserMoneyServiceCases struct {
	MoneyAmountRepo UserMoneyService
}

func NewUserServiceCases() *UserMoneyServiceCases {
	return &UserMoneyServiceCases{
		MoneyAmountRepo: postgres.NewMoneyServiceRepo(),
	}
}

func (u UserMoneyServiceCases) GetUserMoneyAmount(uuid string) (*models.UserMoneyFloat, error) {
	user, err := u.MoneyAmountRepo.GetUserMoneyAmount(uuid)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserMoneyServiceCases) CreateTransaction(request *models.TransactionRequest) (*models.UserMoneyFloat, error) {
	var user *models.UserMoneyFloat
	user, err := u.GetUserMoneyAmount(request.UserUUID)

	if err != nil && err == utils.UserNotFound && request.TransactionType == 1 {
		user, err = u.CreateUser(&models.UserMoneyFloat{UUID: request.UserUUID, Amount: 0})
		if err != nil {
			return nil, err
		}

	} else if err != nil {
		return nil, err
	}

	if request.TransactionType == 0 && user.Amount < request.Amount {
		return nil, utils.NotEnoughFunds
	}

	if request.TransactionType == 1 && user.Amount+request.Amount > 92233720368547758 {
		return nil, utils.NumberIsTooBig
	}

	switch request.TransactionType {
	case 0:
		request.Balance = user.Amount - request.Amount
	case 1:
		request.Balance = user.Amount + request.Amount
	}

	return u.MoneyAmountRepo.CreateTransaction(request)
}

func (u UserMoneyServiceCases) CreateUser(user *models.UserMoneyFloat) (*models.UserMoneyFloat, error) {
	return u.MoneyAmountRepo.CreateUser(user)
}

func (u UserMoneyServiceCases) FundsTransfer(request *models.TransferRequest) (*models.TransferAnswer, error) {
	fromUser, err := u.GetUserMoneyAmount(request.FromUUID)

	if err != nil {
		return nil, err
	}

	if fromUser.Amount < request.Amount {
		return nil, utils.NotEnoughFunds
	}

	_, err = u.GetUserMoneyAmount(request.ToUUID)

	if err != nil && err == utils.UserNotFound {
		_, err = u.CreateUser(&models.UserMoneyFloat{UUID: request.ToUUID, Amount: 0})
		if err != nil {
			return nil, err
		}

	} else if err != nil {
		return nil, err
	}

	return u.MoneyAmountRepo.FundsTransfer(request)
}
