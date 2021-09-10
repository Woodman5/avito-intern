package postgres

import (
	"avito-intern/internal/models"
	"avito-intern/internal/utils"
	"errors"

	"gorm.io/gorm"
)

// func (p *moneyServiceRepo) CreateUser(user *models.User) (*models.User, error) {
// 	result := p.db.Create(&user)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return user, nil
// }

func (p moneyServiceRepo) GetUserMoneyAmount(uuid string) (*models.UserMoney, error) {
	var money models.UserMoney
	result := p.db.First(&money, "uuid = ?", uuid)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, utils.UserNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	return &money, nil
}

// func (p moneyServiceRepo) UpdateUser(user *models.User) (*models.User, error) {
// 	result := p.db.Model(&user).Update("name", user.Name)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	if result.RowsAffected == 0 {
// 		return nil, utils.UserNotFound
// 	}

// 	return user, nil
// }

// func (p moneyServiceRepo) DeleteUser(uuid string) (bool, error) {
// 	result := p.db.Delete(&models.User{}, "uuid = ?", uuid)

// 	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		return false, utils.UserNotFound
// 	} else if result.Error != nil {
// 		return false, result.Error
// 	}

// 	return true, nil
// }

// func (p moneyServiceRepo) GetPetByUUID(userUUID, uuid string) (*models.Pet, error) {
// 	user, err := p.GetUserByUUID(userUUID)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var pet models.Pet

// 	err = p.db.Model(&user).Where("uuid = ?", uuid).Association("Pets").Find(&pet)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if pet.UUID == "" {
// 		return nil, utils.PetNotFound
// 	}

// 	return &pet, nil
// }

// func (p moneyServiceRepo) CreatePet(userUUID string, pet *models.Pet) (*models.Pet, error) {
// 	user, err := p.GetUserByUUID(userUUID)

// 	if err != nil {
// 		return nil, err
// 	}

// 	err = p.db.Model(&user).Association("Pets").Append(pet)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return pet, nil
// }

// func (p moneyServiceRepo) UpdatePet(userUUID string, pet *models.Pet) (*models.Pet, error) {

// 	result := p.db.Model(&pet).Updates(models.Pet{Name: pet.Name, Kind: pet.Kind})

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	if result.RowsAffected == 0 {
// 		return nil, utils.PetNotFound
// 	}

// 	return pet, nil
// }

// func (p moneyServiceRepo) DeletePet(userUUID, uuid string) (bool, error) {
// 	result := p.db.Delete(&models.Pet{}, "uuid = ?", uuid)

// 	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		return false, utils.PetNotFound
// 	} else if result.Error != nil {
// 		return false, result.Error
// 	}

// 	return true, nil
// }
