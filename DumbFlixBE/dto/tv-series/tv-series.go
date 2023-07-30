package tvdto

type CreateTvRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Image       string `json:"image" form:"image" validate:"required"`
	Year        string `json:"year" form:"year" validate:"required"`
	CategoryID  int    `json:"category_id" form:"category_id"`
	Description string `json:"description" form:"description" validate:"required"`
	Link        string `json:"link" form:"link"`
}

type UpdateTvRequest struct {
	Title       string `json:"title" form:"title"`
	Image       string `json:"image" form:"image"`
	Year        string `json:"year" form:"year"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description" form:"description"`
	Link        string `json:"link" form:"link"`
}

type TvResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Year        string `json:"year"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description"`
	Link        string `json:"link"`
}
