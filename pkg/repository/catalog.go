package repository

import (
	"Skipper/pkg/models"
	"encoding/json"
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

//type CatalogElement struct {
//	ID       uint   `json:"id"`
//	Name     string `json:"name"`
//	ParentId uint   `json:"parentId"`
//	Level    uint   `json:"level"`
//}

//type Catalog0 struct {
//	Title string `json:"title"`
//	Id    uint   `json:"index_catalog"`
//	Items []Catalog1
//}
//
//type Catalog1 struct {
//	Index uint   `json:"index1"`
//	Logo  uint   `json:"logo"`
//	Name  string `json:"name"`
//	Items []Catalog2
//}
//
//type Catalog2 struct {
//	Index uint   `json:"index2"`
//	Title string `json:"title"`
//	Items []Catalog3
//}
//
//type Catalog3 struct {
//	Index uint   `json:"index3"`
//	Name  string `json:"name"`
//	Count uint   `json:"count"`
//}

//func (c CatalogPostgres) GetAllCatalog(nodeId uint) string {
//var ids []uint
//var CatalogElements []CatalogElement
//catalogParentsIds := c.db.Raw("SELECT id FROM catalogs WHERE parent_id IS NULL")
//catalogParentsIds.Find(&ids)
//
//for i := 0; i < len(ids); i++ {
//	var results []CatalogElement
//	var parent CatalogElement
//	mainCatalog := c.db.Raw("SELECT id, parent_id, name FROM catalogs WHERE id =?", ids[i])
//	mainCatalog.Find(&parent)
//	data := c.db.Raw("WITH RECURSIVE r (id, parent_id, name, level) AS (SELECT id, parent_id, name, 1 FROM catalogs WHERE parent_id =? UNION ALL SELECT t.id, t.parent_id, t.name, r.level + 1 FROM r INNER JOIN catalogs t ON r.id = t.parent_id) SELECT * FROM r", ids[i])
//	data.Find(&results)
//	CatalogElements = append(CatalogElements, parent)
//	for _, value := range results {
//		CatalogElements = append(CatalogElements, value)
//	}
//}
//var catalog Catalog0
//for i := 0; i < len(CatalogElements); i++ {
//	if CatalogElements[i].Level == 0 {
//		catalog.Title = CatalogElements[i].Name
//		catalog.Id = CatalogElements[i].ID
//		continue
//	} else if CatalogElements[i].Level == 1 {
//		catalog1 := Catalog1{
//			Name:  CatalogElements[i].Name,
//			Logo:  uint(i),
//			Index: CatalogElements[i].ID,
//		}
//		catalog.Items = append(catalog.Items, catalog1)
//	} else if CatalogElements[i].Level == 2 {
//		catalog2 := Catalog2{
//			Title: CatalogElements[i].Name,
//			Index: CatalogElements[i].ID,
//		}
//		catalog.Items[].Items = append(catalog.Items[len(catalog.Items)-1].Items, catalog2)
//	}
//}
//fmt.Println(catalog)
//jsonCatalog, err := json.Marshal(CatalogElements)
//if err != nil {
//	fmt.Printf("Error: %s", err.Error())
//}
//return string(jsonCatalog)
//	return ""
//}

func (c CatalogPostgres) GetAllCatalog(nodeId uint) string {
	//mode := &models.Catalog0{
	//	Name: "Программирование и аналитика",
	//	Child0: []models.Catalog1{
	//		{
	//			Name: "Аналитика",
	//			Child1: []models.Catalog2{
	//				{
	//					Name: "следующий уровень аналитики",
	//					Child2: []models.Catalog3{
	//						{
	//							Name3: "Последний уровень аналитики",
	//							Count: 100,
	//						},
	//					},
	//				},
	//			},
	//		},
	//		{
	//			Name: "Аналитика2",
	//			Child1: []models.Catalog2{
	//				{
	//					Name: "следующий уровень аналитики2",
	//					Child2: []models.Catalog3{
	//						{
	//							Name3: "Последний уровень аналитики 2",
	//							Count: 50,
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//jsonCatalog2, _ := json.Marshal(mode)
	//fmt.Println(string(jsonCatalog2))
	//c.db.Create(mode)

	var m models.Catalog0
	c.db.Preload("Child0.Child1.Child2").Preload(clause.Associations).Find(&m)

	jsonCatalog, _ := json.Marshal(m)
	fmt.Println(string(jsonCatalog))

	return string(jsonCatalog)
}
