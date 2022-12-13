package repository

import (
	"Skipper/pkg/models"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CatalogPostgres struct {
	db *gorm.DB
}

func NewCatalogPostgres(db *gorm.DB) *CatalogPostgres {
	return &CatalogPostgres{db: db}
}

func (c CatalogPostgres) GetCatalog() []models.Catalog0 {
	type Id struct {
		Id string
	}
	var ids []models.Catalog0
	var catalog []models.Catalog0
	c.db.Select("id").Find(&ids)
	for _, d := range ids {
		var m models.Catalog0
		c.db.Preload("Child0.Child1.Child2").Preload(clause.Associations).Find(&m, d.ID)
		catalog = append(catalog, m)
	}
	return catalog
}

type MainCatalog struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (c CatalogPostgres) GetMainCatalog() []MainCatalog {
	var mainCatalogs []MainCatalog
	//var mainCatalogs []models.Catalog1
	c.db.Raw("SELECT id, name FROM catalog1").Find(&mainCatalogs)
	//c.db.Debug().Find(&mainCatalogs)
	return mainCatalogs
}

func (c CatalogPostgres) GetCatalogChild() []models.Catalog3 {
	var catalogChild []models.Catalog3
	c.db.Raw("SELECT id, name3, count FROM catalog3").Find(&catalogChild)
	return catalogChild
}

func (c CatalogPostgres) GetClasses(pagination **models.Pagination) ([]models.User, error) {
	var users []models.User
	offset := ((*pagination).Page - 1) * (*pagination).Limit
	queryBuider := c.db.Limit((*pagination).Limit).Offset(offset)
	filters := fmt.Sprintf("rating >= %d AND rating <= %d AND average_class_price >= %d AND average_class_price <= %d",
		(*pagination).DownRating,
		(*pagination).HighRating,
		(*pagination).DownPrice,
		(*pagination).HighPrice)
	var result *gorm.DB
	if (len((*pagination).Search)) > 0 {
		//queryBuider = queryBuider.Preload("Classes.Tags", "id IN (?)", (*pagination).Search)
		result = queryBuider.
			//Preload("Classes", func(db *gorm.DB) *gorm.DB {
			//	return db.Preload("Tags in (?)", (*pagination).Search)
			//}).
			//Preload("Classes.Tags").
			Preload("Classes").
			Preload("Classes.Tags").
			Preload("Classes.Tags", "id IN (?)", (*pagination).Search).
			Joins("LEFT JOIN (SELECT parent_id as classes_user_id from classes) as classes_data ON classes_data.classes_user_id = id").
			Where("is_mentor = true AND classes_user_id IS NOT NULL AND " + filters).
			Group("\"users\".id").
			Find(&users)
	} else {
		result = queryBuider.
			Preload("Classes").
			Preload("Classes.Tags").
			Joins("LEFT JOIN (SELECT parent_id as classes_user_id from classes) as classes_data ON classes_data.classes_user_id = id").
			Where("is_mentor = true AND classes_user_id IS NOT NULL AND " + filters).
			Group("\"users\".id").
			Find(&users)
	}
	//result = queryBuider.Preload("Classes").
	//	Preload("Classes.Tags").
	//	Joins("LEFT JOIN (SELECT parent_id as classes_user_id from classes) as classes_data ON classes_data.classes_user_id = id").
	//	Where("is_mentor = true AND classes_user_id IS NOT NULL").
	//	Find(&users)

	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return users, nil

}
func (c CatalogPostgres) GetFavouriteMentors(userId uint) ([]models.User, error) {
	var user models.User
	result := c.db.Preload("FavouriteMentors").First(&user, userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return user.FavouriteMentors, nil
}
