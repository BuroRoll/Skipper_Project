package service

import (
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
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
