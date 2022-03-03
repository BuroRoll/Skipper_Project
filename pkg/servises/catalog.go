package service

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
	"encoding/json"
)

type CatalogService struct {
	repo repository.Catalog
}

func NewCatalogService(repo repository.Catalog) *CatalogService {
	return &CatalogService{repo: repo}
}

func (c CatalogService) CreateCatalog(catalog forms.CatalogInput) (uint, error) {
	if catalog.ParentId == nil { //if no parentId (main Catalog)
		mainCatalogId, err := c.repo.CreateMainCatalog(catalog.Name)
		if err != nil {
			return 0, err
		}
		return mainCatalogId, nil
	} else {
		childId, err := c.repo.CreateChildCatalog(catalog.Name, catalog.ParentId)
		if err != nil {
			return 0, err
		}
		return childId, nil
	}
}

func (c CatalogService) GetCatalog() string {
	data := c.repo.GetCatalog()
	return data
}

func (c CatalogService) GetMainCatalog() string {
	data := c.repo.GetMainCatalog()
	mainCatalog, _ := json.Marshal(data)
	return string(mainCatalog)
}

func (c CatalogService) GetCatalogChild() (string, error) {
	catalogChild := c.repo.GetCatalogChild()
	jsonData, err := json.Marshal(catalogChild)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

type ClassesData struct {
	Id             uint   `json:"ID"`
	FirstName      string `json:"FirstName"`
	SecondName     string `json:"SecondName"`
	ProfilePicture string `json:"ProfilePicture"`
	Description    string `json:"Description"`
	Specialization string `json:"Specialization"`
	//Classes        []models.Class `json:"classes"`
	Classes []C `json:"classes"`
}

type C struct {
	ID          uint   `json:"ID"`
	ClassName   string `json:"ClassName"`
	Description string `json:"Description"`
	Tags        []T    `json:"Tags"`
}

type T struct {
	ID   uint   `json:"ID"`
	Name string `json:"name3"`
}

func (c CatalogService) GetClasses(pagination *models.Pagination) (string, error) {

	data, err := c.repo.GetClasses(&pagination)
	if err != nil {
		return "", nil
	}
	j, _ := json.Marshal(data)
	var d []ClassesData
	err = json.Unmarshal(j, &d)
	if err != nil {
		return "", err
	}
	p, _ := json.Marshal(d)

	return string(p), nil
}
