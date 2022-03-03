package repository

import (
	"Skipper/pkg/models"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CatalogPostgres struct {
	db *gorm.DB
}

func NewCatalogPostgres(db *gorm.DB) *CatalogPostgres {
	return &CatalogPostgres{db: db}
}

func (c CatalogPostgres) CreateMainCatalog(catalogName string) (uint, error) {
	return 0, nil
}

func (c CatalogPostgres) CreateChildCatalog(catalogName string, parentid *uint) (uint, error) {
	return 0, nil
}

func (c CatalogPostgres) GetCatalog() string {
	type Id struct {
		Id string
	}
	var ids []Id
	var catalog []models.Catalog0
	c.db.Raw("SELECT id FROM catalog0").Find(&ids)
	for _, d := range ids {
		var m models.Catalog0
		c.db.Preload("Child0.Child1.Child2").Preload(clause.Associations).Find(&m, d.Id)
		catalog = append(catalog, m)
	}
	jsonCatalog, _ := json.Marshal(catalog)
	return string(jsonCatalog)
}

type MainCatalog struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (c CatalogPostgres) GetMainCatalog() []MainCatalog {
	var mainCatalogs []MainCatalog
	c.db.Raw("SELECT id, name FROM catalog1").Find(&mainCatalogs)
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
	queryBuider := c.db.Debug().Limit((*pagination).Limit).Offset(offset).Order((*pagination).Sort)
	var result *gorm.DB
	if (len((*pagination).Search)) > 0 {
		result = queryBuider.Debug().Preload("Classes").Preload("Classes.Tags").Preload("Classes.Tags", "id IN (?)", (*pagination).Search).Where("is_mentor = true").Find(&users)
	} else {
		result = queryBuider.Debug().Preload("Classes").Preload("Classes.Tags").Where("is_mentor = true").Find(&users)
	}
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	//re := Filter(users, func(user models.User) bool {
	//	return len(user.Classes) > 0
	//})
	//return re, nil
	return users, nil

}

func Filter(arr []models.User, cond func(user models.User) bool) []models.User {
	var result []models.User
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}
