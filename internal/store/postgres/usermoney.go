package postgres

import (
	"avito-intern/internal/models"
	"avito-intern/internal/utils"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math"
)

func (p *moneyServiceRepo) CreateTransaction(request *models.TransactionRequest) (*models.UserMoneyFloat, error) {
	intBalance := uint64(math.Floor(request.Balance * 100))

	var money = &models.UserMoney{UUID: request.UserUUID, Amount: intBalance}
	result := p.db.Model(&money).Update("amount", money.Amount)

	if result.Error != nil {
		return nil, result.Error
	}

	trUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	var tr = &models.Transaction{
		UUID:            trUUID.String(),
		UserUUID:        request.UserUUID,
		TransactionType: request.TransactionType,
		Amount:          uint64(math.Floor(request.Amount * 100)),
		Balance:         intBalance,
		Source:          request.Source,
		Reason:          request.Reason,
	}

	result = p.db.Create(&tr)

	if result.Error != nil {
		return nil, result.Error
	}

	floatMoney := float64(money.Amount) / 100

	return &models.UserMoneyFloat{UUID: money.UUID, Amount: floatMoney}, nil
}

func (p moneyServiceRepo) GetUserMoneyAmount(uuid string) (*models.UserMoneyFloat, error) {
	var money models.UserMoney
	result := p.db.First(&money, "uuid = ?", uuid)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, utils.UserNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	floatMoney := float64(money.Amount) / 100

	return &models.UserMoneyFloat{UUID: money.UUID, Amount: floatMoney}, nil
}

func (p moneyServiceRepo) CreateUser(user *models.UserMoneyFloat) (*models.UserMoneyFloat, error) {
	var money = &models.UserMoney{UUID: user.UUID}
	money.Amount = uint64(math.Floor(user.Amount * 100))

	if money.Amount > 9223372036854775800 {
		return nil, utils.NumberIsTooBig
	}

	result := p.db.Create(&money)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (p *moneyServiceRepo) FundsTransfer(request *models.TransferRequest) (*models.TransferAnswer, error) {
	var fromUser models.UserMoney
	var toUser models.UserMoney

	intAmount := uint64(math.Floor(request.Amount * 100))

	err := p.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.First(&fromUser, "uuid = ?", request.FromUUID).Error; err != nil {
			return err
		}
		if err := tx.First(&toUser, "uuid = ?", request.ToUUID).Error; err != nil {
			return err
		}

		if toUser.Amount+intAmount > 9223372036854775800 {
			return utils.NumberIsTooBig
		}

		newFromBalance := fromUser.Amount - intAmount
		newToBalance := toUser.Amount + intAmount

		if err := tx.Model(&fromUser).Update("amount", newFromBalance).Error; err != nil {
			return err
		}

		if err := tx.Model(&toUser).Update("amount", newToBalance).Error; err != nil {
			return err
		}

		fromUser.Amount = newFromBalance
		toUser.Amount = newToBalance

		return nil
	})

	if err != nil {
		return nil, err
	}

	fromUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	toUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	var trFrom = &models.Transaction{
		UUID:            fromUUID.String(),
		UserUUID:        fromUser.UUID,
		TransactionType: 0,
		Amount:          intAmount,
		Balance:         fromUser.Amount,
		Source:          "funds transfer",
		Reason:          fmt.Sprintf("transfer to user: %s", toUser.UUID),
	}

	var trTo = &models.Transaction{
		UUID:            toUUID.String(),
		UserUUID:        toUser.UUID,
		TransactionType: 1,
		Amount:          intAmount,
		Balance:         toUser.Amount,
		Source:          "funds transfer",
		Reason:          fmt.Sprintf("transfer from user: %s", fromUser.UUID),
	}

	result := p.db.Create(&trFrom)

	if result.Error != nil {
		return nil, result.Error
	}

	result = p.db.Create(&trTo)

	if result.Error != nil {
		return nil, result.Error
	}

	floatFromMoney := float64(fromUser.Amount) / 100
	floatToMoney := float64(toUser.Amount) / 100

	resFrom := models.UserMoneyFloat{UUID: fromUser.UUID, Amount: floatFromMoney}
	resTo := models.UserMoneyFloat{UUID: toUser.UUID, Amount: floatToMoney}

	return &models.TransferAnswer{To: resTo, From: resFrom}, nil
}
