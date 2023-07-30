package tvepisodedto

type CreateEpisode struct {
	Title string `json:"title" form:"title" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
	Link  string `json:"link" form:"link" validate:"required"`
	TvID  int    `json:"tv_id" form:"tv_id"`
}

type UpdateEpisode struct {
	Title string `json:"title" form:"title"`
	Image string `json:"image" form:"image"`
	Link  string `json:"link" form:"link"`
	TvID  int    `json:"tv_id"`
}

type EpisodeTVResponse struct {
	Title string `json:"title"`
	Image string `json:"image"`
	Link  string `json:"link"`
	TvID  int    `json:"tv_id"`
}
