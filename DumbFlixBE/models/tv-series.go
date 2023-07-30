package models

type Tv struct {
	ID          int              `json:"id" gorm:"primaryKey:autoIncrement"`
	Title       string           `json:"title" gorm:"type: varchar(255)"`
	Image       string           `json:"image" gorm:"type: varchar(255)"`
	Year        string           `json:"year" gorm:"type: varchar(255)"`
	CategoryID  int              `json:"category_id"`
	Category    CategoryResponse `json:"category"`
	Description string           `json:"description" gorm:"type: varchar(255)"`
	Link        string           `json:"link" gorm:"type: varchar(255)"`
}

type TvResponse struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	Image       string           `json:"image"`
	Year        string           `json:"year"`
	CategoryID  int              `json:"category_id"`
	Category    CategoryResponse `json:"category"`
	Description string           `json:"description"`
	Link        string           `json:"link"`
}

func (TvResponse) TableName() string {
	return "tvs"
}
