package models

type TypeFilm struct {
	ID   int    `json:"id" gorm:"primaryKey:autoIncrement"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}

type TypeFilmResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (TypeFilmResponse) TableName() string {
	return "types"
}
