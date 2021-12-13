package repository

import (
	"Skipper/pkg/models"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClassesPostgres struct {
	db *gorm.DB
}

func NewClassesPostgres(db *gorm.DB) *ClassesPostgres {
	return &ClassesPostgres{db: db}
}

func (c ClassesPostgres) CreateUserClasses(class models.Class) (uint, error) {
	result := c.db.Create(&class)
	if errors.Is(result.Error, gorm.ErrRegistered) {
		return class.ID, result.Error
	}
	return class.ID, nil
}

func (c ClassesPostgres) CreateTheoreticClass(theoreticClass models.TheoreticClass) error {
	result := c.db.Create(&theoreticClass)
	if errors.Is(result.Error, gorm.ErrRegistered) {
		return result.Error
	}
	return nil
}

func (c ClassesPostgres) CreatePracticClass(practicClass models.PracticClass) error {
	result := c.db.Create(&practicClass)
	if errors.Is(result.Error, gorm.ErrRegistered) {
		return result.Error
	}
	return nil
}

func (c ClassesPostgres) CreateKeyClass(keyClass models.KeyClass) error {
	result := c.db.Create(&keyClass)
	if errors.Is(result.Error, gorm.ErrRegistered) {
		return result.Error
	}
	return nil
}

func (c ClassesPostgres) GetCatalogTags(catalogId uint) (models.Catalog3, error) {
	var tag models.Catalog3
	result := c.db.First(&tag, catalogId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return tag, result.Error
	}
	return tag, nil
}

func (c ClassesPostgres) GetUserClasses(userId uint) ([]models.Class, error) {
	var classes []models.Class
	result := c.db.Raw("SELECT * FROM classes WHERE parent_id = ?", userId).Preload("Tags").Preload(clause.Associations).Find(&classes)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return classes, nil
}
