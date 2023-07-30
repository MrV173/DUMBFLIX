package models

type EpisodeTV struct {
	ID    int        `json:"id" gorm:"primaryKey:autoIncrement"`
	Title string     `json:"title" gorm:"type: varchar(255)"`
	Image string     `json:"image" gorm:"type: varchar(255)"`
	Link  string     `json:"link" gorm:"type: varchar(255)"`
	TvID  int        `json:"tv_id"`
	Tv    TvResponse `json:"tv"`
}

type EpisodeTVResponse struct {
	ID    int        `json:"id"`
	Title string     `json:"title"`
	Image string     `json:"image"`
	Link  string     `json:"link"`
	TvID  int        `json:"tv_id"`
	Tv    TvResponse `json:"tv"`
}

func (EpisodeTVResponse) TableName() string {
	return "episodetvs"
}
