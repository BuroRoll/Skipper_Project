package forms

type BookingClassInput struct {
	ClassType string `json:"class_type" binding:"required"`
	ClassId   uint   `json:"class_id" binding:"required"`

	MentorId uint `json:"mentor_id" binding:"required"`

	Duration15 bool `json:"duration_15"`
	Price15    uint `json:"price_15"`

	Duration30_1 bool `json:"duration_30_1"`
	Price30_1    uint `json:"price_30_1"`
	Duration30_3 bool `json:"duration_30_3"`
	Price30_3    uint `json:"price_30_3"`
	Duration30_5 bool `json:"duration_30_5"`
	Price30_5    uint `json:"price_30_5"`

	Duration60_1 bool `json:"duration_60_1"`
	Price60_1    uint `json:"price_60_1"`
	Duration60_3 bool `json:"duration_60_3"`
	Price60_3    uint `json:"price_60_3"`
	Duration60_5 bool `json:"duration_60_5"`
	Price60_5    uint `json:"price_60_5"`

	Duration90_1 bool `json:"duration_90_1"`
	Price90_1    uint `json:"price_90_1"`
	Duration90_3 bool `json:"duration_90_3"`
	Price90_3    uint `json:"price_90_3"`
	Duration90_5 bool `json:"duration_90_5"`
	Price90_5    uint `json:"price_90_5"`

	FullTime      bool `json:"full_time"`
	PriceFullTime uint `json:"price_full_time"`

	Time []string `json:"time" binding:"required"`

	Communication uint `json:"communication" binding:"required"`
}
