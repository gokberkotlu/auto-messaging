package repository

import (
	"fmt"

	"github.com/gokberkotlu/auto-messaging/database"
	"github.com/gokberkotlu/auto-messaging/entity"
	"gorm.io/gorm"
)

type IMigrationProcessRepository interface {
	GetFirst() (entity.MigrationProcess, error)
	Create(status entity.MigrationProcessStatusType) error
	Update(status entity.MigrationProcessStatusType) error
}

type MigrationProcessRepository struct {
	db *gorm.DB
}

func NewMigrationProcessRepository() IMigrationProcessRepository {
	db, err := database.GetDB()
	if err != nil {
		fmt.Errorf("failed to create migration process repository: %w", err)
	}
	return &MigrationProcessRepository{
		db: db,
	}
}

func (repository *MigrationProcessRepository) GetFirst() (entity.MigrationProcess, error) {
	var migrationProcess entity.MigrationProcess

	err := database.CheckIfDbConnectionInitialized()
	if err != nil {
		return migrationProcess, err
	}

	if err := repository.db.Order("id").First(&migrationProcess).Error; err != nil {
		return migrationProcess, err
	}

	return migrationProcess, nil
}

func (repository *MigrationProcessRepository) Create(status entity.MigrationProcessStatusType) error {
	err := database.CheckIfDbConnectionInitialized()
	if err != nil {
		return err
	}

	if err := repository.db.Create(&entity.MigrationProcess{}).Error; err != nil {
		return err
	}

	return nil
}

func (repository *MigrationProcessRepository) Update(status entity.MigrationProcessStatusType) error {
	err := database.CheckIfDbConnectionInitialized()
	if err != nil {
		return err
	}

	migrationProcess, err := repository.GetFirst()
	if err != nil {
		return err
	}

	if err = repository.db.Model(&migrationProcess).Update("status", status).Error; err != nil {
		return fmt.Errorf(`"update migration process status" query failed: %s`, err.Error())
	}

	return nil
}
