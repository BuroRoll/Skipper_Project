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

func (c ClassesPostgres) CreateTheoreticClass(theoreticClass models.TheoreticClass) (uint, error) {
	result := c.db.Create(&theoreticClass)
	if errors.Is(result.Error, gorm.ErrRegistered) {
		return 0, result.Error
	}
	return theoreticClass.ID, nil
}

func (c ClassesPostgres) CreatePracticClass(practicClass models.PracticClass) (uint, error) {
	result := c.db.Create(&practicClass)
	if errors.Is(result.Error, gorm.ErrRegistered) {
		return 0, result.Error
	}
	return practicClass.ID, nil
}

func (c ClassesPostgres) CreateKeyClass(keyClass models.KeyClass) (uint, error) {
	result := c.db.Create(&keyClass)
	if errors.Is(result.Error, gorm.ErrRegistered) {
		return 0, result.Error
	}
	return keyClass.ID, nil
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
	result := c.db.Raw("SELECT * FROM classes WHERE parent_id = ? AND deleted_at IS NULL", userId).Preload("Tags").Preload(clause.Associations).Find(&classes)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return classes, nil
}

func (c ClassesPostgres) DeleteClass(classId string) error {
	//var class models.Class
	//c.db.First(&class, classId)

	c.db.Exec("DELETE FROM catalog3_class WHERE class_id = ?;", classId)
	c.db.Exec("DELETE FROM practic_classes WHERE class_parent_id = ?;", classId)
	c.db.Exec("DELETE FROM theoretic_classes WHERE class_parent_id = ?;", classId)
	c.db.Exec("DELETE FROM key_classes WHERE class_parent_id = ?;", classId)
	c.db.Exec("DELETE FROM classes WHERE id = ?", classId)
	return nil
}

func (c ClassesPostgres) DeleteTheoreticClass(classId string) error {
	result := c.db.Unscoped().Delete(&models.TheoreticClass{}, classId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (c ClassesPostgres) DeletePracticClass(classId string) error {
	result := c.db.Unscoped().Delete(&models.PracticClass{}, classId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (c ClassesPostgres) DeleteKeyClass(classId string) error {
	result := c.db.Unscoped().Delete(&models.KeyClass{}, classId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (c ClassesPostgres) UpdateClass(classData models.Class, classId uint) error {
	var class models.Class
	c.db.First(&class, classId)
	class.ClassName = classData.ClassName
	class.Description = classData.Description
	c.db.Exec("DELETE FROM catalog3_class WHERE class_id = ?", classId)
	class.Tags = classData.Tags
	result := c.db.Save(&class)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (c ClassesPostgres) UpdateTheoreticClass(classData models.TheoreticClass, classId uint) error {
	var theoreticClass models.TheoreticClass
	c.db.First(&theoreticClass, classId)
	theoreticClass.Duration15 = classData.Duration15
	theoreticClass.Duration30 = classData.Duration30
	theoreticClass.Duration60 = classData.Duration60
	theoreticClass.Duration90 = classData.Duration90
	theoreticClass.Price15 = classData.Price15
	theoreticClass.Price30 = classData.Price30
	theoreticClass.Price60 = classData.Price60
	theoreticClass.Price90 = classData.Price90
	theoreticClass.Time = classData.Time
	result := c.db.Save(&theoreticClass)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (c ClassesPostgres) UpdatePracticClass(classData models.PracticClass, classId uint) error {
	var practicClass models.PracticClass
	c.db.First(&practicClass, classId)
	practicClass.Duration15 = classData.Duration15
	practicClass.Duration30 = classData.Duration30
	practicClass.Duration60 = classData.Duration60
	practicClass.Duration90 = classData.Duration90
	practicClass.Price15 = classData.Price15
	practicClass.Price30 = classData.Price30
	practicClass.Price60 = classData.Price60
	practicClass.Price90 = classData.Price90
	practicClass.Time = classData.Time
	result := c.db.Save(&practicClass)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (c ClassesPostgres) UpdateKeyClass(classData models.KeyClass, classId uint) error {
	var keyClass models.KeyClass
	c.db.First(&keyClass, classId)
	keyClass.Duration15 = classData.Duration15
	keyClass.Price15 = classData.Price15
	keyClass.FullTime = classData.FullTime
	keyClass.PriceFullTime = classData.PriceFullTime
	keyClass.Time = classData.Time
	result := c.db.Save(&keyClass)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (c ClassesPostgres) GetClassById(classId string) (models.Class, error) {
	var class models.Class
	result := c.db.Preload(clause.Associations).First(&class, classId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Class{}, result.Error
	}
	return class, nil
}
