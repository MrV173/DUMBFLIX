package typedto

type CreateTypeFilmRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type TypeFilmResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
