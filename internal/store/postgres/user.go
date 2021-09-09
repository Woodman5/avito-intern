package postgres

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"simbirsoft/gohw/hw6/internal/app/hispet/models"
	"simbirsoft/gohw/hw6/internal/app/hispet/utils"
	"time"
)

func getDBDSN() string {
	dbHost, ok := os.LookupEnv("DATABASE_HOST")
	if !ok {
		dbHost = "localhost"
	}

	dbUser, ok := os.LookupEnv("DATABASE_USER")
	if !ok {
		dbUser = "postgres"
	}

	dbPassword, ok := os.LookupEnv("DATABASE_PASSWORD")
	if !ok {
		dbPassword = "superSecretWord"
	}

	dbPort, ok := os.LookupEnv("DATABASE_PORT")
	if !ok {
		dbPort = "5432"
	}

	dbName, ok := os.LookupEnv("DATABASE_NAME")
	if !ok {
		dbName = "postgres"
	}

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Europe/Moscow"

	return dsn
}

type userServiceRepo struct {
	db *gorm.DB
}

func NewUserServiceRepo() *userServiceRepo {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(getDBDSN()), &gorm.Config{Logger: newLogger})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &userServiceRepo{db: db}
}

func (p *userServiceRepo) CreateUser(user *models.User) (*models.User, error) {
	result := p.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (p userServiceRepo) GetUserByUUID(uuid string) (*models.User, error) {
	var user models.User
	result := p.db.First(&user, "uuid = ?", uuid)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, utils.UserNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (p userServiceRepo) UpdateUser(user *models.User) (*models.User, error) {
	result := p.db.Model(&user).Update("name", user.Name)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, utils.UserNotFound
	}

	return user, nil
}

func (p userServiceRepo) DeleteUser(uuid string) (bool, error) {
	result := p.db.Delete(&models.User{}, "uuid = ?", uuid)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, utils.UserNotFound
	} else if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (p userServiceRepo) GetPetByUUID(userUUID, uuid string) (*models.Pet, error) {
	user, err := p.GetUserByUUID(userUUID)

	if err != nil {
		return nil, err
	}

	var pet models.Pet

	err = p.db.Model(&user).Where("uuid = ?", uuid).Association("Pets").Find(&pet)

	if err != nil {
		return nil, err
	}

	if pet.UUID == "" {
		return nil, utils.PetNotFound
	}

	return &pet, nil
}

func (p userServiceRepo) CreatePet(userUUID string, pet *models.Pet) (*models.Pet, error) {
	user, err := p.GetUserByUUID(userUUID)

	if err != nil {
		return nil, err
	}

	err = p.db.Model(&user).Association("Pets").Append(pet)

	if err != nil {
		return nil, err
	}

	return pet, nil
}

func (p userServiceRepo) UpdatePet(userUUID string, pet *models.Pet) (*models.Pet, error) {

	result := p.db.Model(&pet).Updates(models.Pet{Name: pet.Name, Kind: pet.Kind})

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, utils.PetNotFound
	}

	return pet, nil
}

func (p userServiceRepo) DeletePet(userUUID, uuid string) (bool, error) {
	result := p.db.Delete(&models.Pet{}, "uuid = ?", uuid)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, utils.PetNotFound
	} else if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
