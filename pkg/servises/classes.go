package service

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
	"encoding/json"
)

type ClassesService struct {
	repo repository.Classes
}

func NewClassesService(repo repository.Classes) *ClassesService {
	return &ClassesService{repo: repo}
}

func (c ClassesService) CreateUserClass(class forms.ClassesInput, userId uint) (uint, error) {
	var classBd models.Class
	classBd = models.Class{
		ParentId:    userId,
		ClassName:   class.ClassName,
		Description: class.Description,
	}
	for _, id := range class.Tags {
		tag, err := c.repo.GetCatalogTags(id)
		if err != nil {
			return 1, err
		}
		classBd.Tags = append(classBd.Tags, &tag)
	}
	return c.repo.CreateUserClasses(classBd)
}

func (c ClassesService) CreateTheoreticClass(class forms.TheoreticClassInput, userId uint) error {
	var classBd models.TheoreticClass
	classBd = models.TheoreticClass{
		ClassParentId: class.ParentId,

		Duration15: class.Duration15,
		Price15:    class.Price15,
		Duration30: class.Duration30,
		Price30:    class.Price30,
		Duration60: class.Duration60,
		Price60:    class.Price60,
		Duration90: class.Duration90,
		Price90:    class.Price90,

		Time: class.Time,
	}
	return c.repo.CreateTheoreticClass(classBd)
}

func (c ClassesService) CreatePracticClass(class forms.PracticClassInput, userId uint) error {
	var classBd models.PracticClass
	classBd = models.PracticClass{
		ClassParentId: class.ParentId,

		Duration15: class.Duration15,
		Price15:    class.Price15,
		Duration30: class.Duration30,
		Price30:    class.Price30,
		Duration60: class.Duration60,
		Price60:    class.Price60,
		Duration90: class.Duration90,
		Price90:    class.Price90,

		Time: class.Time,
	}
	return c.repo.CreatePracticClass(classBd)
}

func (c ClassesService) CreateKeyClass(class forms.KeyClass, userId uint) error {
	var classBd models.KeyClass
	classBd = models.KeyClass{
		ClassParentId: class.ParentId,

		Duration15:    class.Duration15,
		Price15:       class.Price15,
		FullTime:      class.FullTime,
		PriceFullTime: class.PriceFullTime,

		Time: class.Time,
	}
	return c.repo.CreateKeyClass(classBd)
}

func (c ClassesService) GetUserClasses(userId uint) (string, error) {
	classes, err := c.repo.GetUserClasses(userId)
	if err != nil {
		return "", err
	}
	jsonData, err := json.Marshal(classes)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
