package handlers

import (
	dto "dumbflix/dto/result"
	typedto "dumbflix/dto/type"
	"dumbflix/models"
	"dumbflix/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HandleTypeFilm struct {
	TypeFilmRepository repository.TypeFilmRepository
}

func TypeFilmHandler(typeFilmRepository repository.TypeFilmRepository) *HandleTypeFilm {
	return &HandleTypeFilm{typeFilmRepository}
}

func (h *HandleTypeFilm) FindTypeFilms(c echo.Context) error {
	typefilms, err := h.TypeFilmRepository.FindTypeFilms()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: typefilms,
	})
}

func (h *HandleTypeFilm) GetTypeFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	typeFilm, err := h.TypeFilmRepository.GetTypeFilm(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertTypeFilmResponse(typeFilm)})
}

func (h *HandleTypeFilm) CreateTypeFilm(c echo.Context) error {
	request := new(typedto.CreateTypeFilmRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	data := models.TypeFilm{
		Name: request.Name,
	}

	response, err := h.TypeFilmRepository.CreateTypeFilm(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})

}

func (h *HandleTypeFilm) DeleteTypeFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	typeFilm, err := h.TypeFilmRepository.GetTypeFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TypeFilmRepository.DeleteTypeFilm(typeFilm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertTypeFilmResponse(data),
	})
}

func convertTypeFilmResponse(u models.TypeFilm) typedto.TypeFilmResponse {
	return typedto.TypeFilmResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
