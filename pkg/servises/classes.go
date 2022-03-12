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

func (c ClassesService) CreateTheoreticClass(class forms.TheoreticClassInput, userId uint) (uint, error) {
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

func (c ClassesService) CreatePracticClass(class forms.PracticClassInput, userId uint) (uint, error) {
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

func (c ClassesService) CreateKeyClass(class forms.KeyClass, userId uint) (uint, error) {
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

func (c ClassesService) DeleteClass(classId string) error {
	return c.repo.DeleteClass(classId)
}

func (c ClassesService) DeleteTheoreticClass(classId string) error {
	return c.repo.DeleteTheoreticClass(classId)
}

func (c ClassesService) DeletePracticClass(classId string) error {
	return c.repo.DeletePracticClass(classId)
}

func (c ClassesService) DeleteKeyClass(classId string) error {
	return c.repo.DeleteKeyClass(classId)
}

func (c ClassesService) UpdateClass(classData forms.UpdateClassesInput) error {
	var classDB models.Class
	classDB = models.Class{
		ClassName:   classData.ClassName,
		Description: classData.Description,
	}
	for _, id := range classData.Tags {
		tag, err := c.repo.GetCatalogTags(id)
		if err != nil {
			return err
		}
		classDB.Tags = append(classDB.Tags, &tag)
	}
	return c.repo.UpdateClass(classDB, classData.ClassId)
}

func (c ClassesService) UpdateTheoreticClass(theoreticClassData forms.UpdateSubclassInput) error {
	theoreticClassDB := models.TheoreticClass{
		Duration15: theoreticClassData.Duration15,
		Price15:    theoreticClassData.Price15,
		Duration30: theoreticClassData.Duration30,
		Price30:    theoreticClassData.Price30,
		Duration60: theoreticClassData.Duration60,
		Price60:    theoreticClassData.Price60,
		Duration90: theoreticClassData.Duration90,
		Price90:    theoreticClassData.Price90,

		Time: theoreticClassData.Time,
	}
	return c.repo.UpdateTheoreticClass(theoreticClassDB, theoreticClassData.ClassId)
}

func (c ClassesService) UpdatePracticClass(practicClassData forms.UpdateSubclassInput) error {
	practicClassDB := models.PracticClass{
		Duration15: practicClassData.Duration15,
		Price15:    practicClassData.Price15,
		Duration30: practicClassData.Duration30,
		Price30:    practicClassData.Price30,
		Duration60: practicClassData.Duration60,
		Price60:    practicClassData.Price60,
		Duration90: practicClassData.Duration90,
		Price90:    practicClassData.Price90,

		Time: practicClassData.Time,
	}
	return c.repo.UpdatePracticClass(practicClassDB, practicClassData.ClassId)
}

func (c ClassesService) UpdateKeyClass(keyClassData forms.UpdateKeyClassInput) error {

	keyClassDB := models.KeyClass{
		Duration15:    keyClassData.Duration15,
		Price15:       keyClassData.Price15,
		FullTime:      keyClassData.FullTime,
		PriceFullTime: keyClassData.PriceFullTime,
		Time:          keyClassData.Time,
	}
	return c.repo.UpdateKeyClass(keyClassDB, keyClassData.ClassId)
}

type TPClass struct {
	Id            uint `json:"ID"`
	ClassParentId uint `json:"ClassParentId"`
	Duration15    bool `json:"Duration15"`
	Price15       uint `json:"Price15"`

	Duration30 bool `json:"Duration30"`
	Price30    uint `json:"Price30"`

	Duration60 bool `json:"Duration60"`
	Price60    uint `json:"Price60"`

	Duration90 bool `json:"Duration90"`
	Price90    uint `json:"Price90"`

	Time string `json:"Time"`
}

type KClass struct {
	Id            uint `json:"ID"`
	ClassParentId uint `json:"ClassParentId"`

	Duration15 bool `json:"Duration15"`
	Price15    uint `json:"Price15"`

	FullTime      bool `json:"FullTime"`
	PriceFullTime uint `json:"PriceFullTime"`

	Time string `json:"Time"`
}

func (c ClassesService) GetClassById(classId string) (string, string, string, error) {
	class, err := c.repo.GetClassById(classId)
	if err != nil {
		return "", "", "", err
	}
	var TheoreticClassData TPClass
	var PracticClassData TPClass
	var KeyClassData KClass
	tc, _ := json.Marshal(class.TheoreticClass)
	err = json.Unmarshal(tc, &TheoreticClassData)
	pc, _ := json.Marshal(class.PracticClass)
	err = json.Unmarshal(pc, &PracticClassData)
	kc, _ := json.Marshal(class.KeyClass)
	err = json.Unmarshal(kc, &KeyClassData)

	tc, _ = json.Marshal(TheoreticClassData)
	pc, _ = json.Marshal(PracticClassData)
	kc, _ = json.Marshal(KeyClassData)
	return string(tc), string(pc), string(kc), nil
}
