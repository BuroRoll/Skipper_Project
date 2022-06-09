package forms

type ClassesInput struct {
	ClassName   string `json:"class_name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Tags        []uint `json:"tags" binding:"required"`
}

type UpdateClassesInput struct {
	ClassId     uint   `json:"class_id" binding:"required"`
	ClassName   string `json:"class_name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Tags        []uint `json:"tags" binding:"required"`
}

type UpdateSubclassInput struct {
	ClassId uint `json:"class_id" binding:"required"`

	Duration15 bool `json:"duration_15"`
	Price15    uint `json:"price_15"`

	Duration30 bool `json:"duration_30"`
	Price30    uint `json:"price_30"`

	Duration60 bool `json:"duration_60"`
	Price60    uint `json:"price_60"`

	Duration90 bool `json:"duration_90"`
	Price90    uint `json:"price_90"`

	Time string `json:"time" binding:"required"`
}

type UpdateKeyClassInput struct {
	ClassId uint `json:"class_id" binding:"required"`

	Duration15 bool `json:"duration_15"`
	Price15    uint `json:"price_15"`

	FullTime      bool `json:"full_time"`
	PriceFullTime uint `json:"price_full_time"`

	Time string `json:"time" binding:"required"`
}

type TheoreticClassInput struct {
	ParentId uint `json:"parent_id" binding:"required"`

	Duration15 bool `json:"duration_15"`
	Price15    uint `json:"price_15"`

	Duration30 bool `json:"duration_30"`
	Price30    uint `json:"price_30"`

	Duration60 bool `json:"duration_60"`
	Price60    uint `json:"price_60"`

	Duration90 bool `json:"duration_90"`
	Price90    uint `json:"price_90"`

	Time string `json:"time" binding:"required"`
}

type PracticClassInput struct {
	ParentId uint `json:"parent_id" binding:"required"`

	Duration15 bool `json:"duration_15"`
	Price15    uint `json:"price_15"`

	Duration30 bool `json:"duration_30"`
	Price30    uint `json:"price_30"`

	Duration60 bool `json:"duration_60"`
	Price60    uint `json:"price_60"`

	Duration90 bool `json:"duration_90"`
	Price90    uint `json:"price_90"`

	Time string `json:"time" binding:"required"`
}

type KeyClass struct {
	ParentId uint `json:"parent_id" binding:"required"`

	Duration15 bool `json:"duration_15"`
	Price15    uint `json:"price_15"`

	FullTime      bool `json:"full_time"`
	PriceFullTime uint `json:"price_full_time"`

	Time string `json:"time" binding:"required"`
}
