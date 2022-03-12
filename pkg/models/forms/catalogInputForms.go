package forms

type CatalogInput struct {
	Name     string `json:"name" binding:"required"`
	ParentId *uint  `json:"parent_id"`
}
