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
	//var catalog models.Catalog
	//catalog = models.Catalog{
	//	Name:     catalogName,
	//	ParentId: nil,
	//}
	//result := c.db.Create(&catalog)
	//if result.Error != nil {
	//	return catalog.ID, result.Error
	//}
	//return catalog.ID, nil
	return 0, nil
}

func (c CatalogPostgres) CreateChildCatalog(catalogName string, parentid *uint) (uint, error) {
	//var catalog models.Catalog
	//catalog = models.Catalog{
	//	Name:     catalogName,
	//	ParentId: parentid,
	//}
	//result := c.db.Create(&catalog)
	//if result.Error != nil {
	//	return catalog.ID, result.Error
	//}
	//return catalog.ID, nil
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
