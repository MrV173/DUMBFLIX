package handlers

import (
	dto "dumbflix/dto/result"
	tvdto "dumbflix/dto/tv-series"
	"dumbflix/models"
	"dumbflix/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_cover = "http://localhost:5000/uploads"

type HandlerTv struct {
	TvRepository repository.TvRepository
}

func TvHandler(tvRepository repository.TvRepository) *HandlerTv {
	return &HandlerTv{tvRepository}
}

func (h *HandlerTv) FindTvs(c echo.Context) error {
	tvs, err := h.TvRepository.FindTvs()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	for i, p := range tvs {
		tvs[i].Image = path_cover + p.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: tvs,
	})
}

func (h *HandlerTv) GetTv(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tv, err := h.TvRepository.GetTv(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	tv.Image = path_cover + tv.Image

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertTvResponse(tv)})
}

func (h *HandlerTv) CreateTv(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	category_id, _ := strconv.Atoi(c.FormValue("category_id"))
	fmt.Println("ini data tv", dataFile)

	request := tvdto.CreateTvRequest{
		Title:       c.FormValue("title"),
		Image:       dataFile,
		Year:        c.FormValue("year"),
		CategoryID:  category_id,
		Description: c.FormValue("description"),
		Link:        c.FormValue("link"),
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	data := models.Tv{
		Title:       request.Title,
		Image:       request.Image,
		Year:        request.Year,
		CategoryID:  request.CategoryID,
		Description: request.Description,
		Link:        request.Link,
	}

	response, err := h.TvRepository.CreateTv(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func (h *HandlerTv) DeleteTv(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tv, err := h.TvRepository.GetTv(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TvRepository.DeleteTv(tv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertTvResponse(data),
	})
}

func (h *HandlerTv) UpdateTv(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tv, err := h.TvRepository.GetTv(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	request := new(tvdto.UpdateTvRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if request.Title != "" {
		tv.Title = request.Title
	}

	if request.Image != "" {
		tv.Image = request.Image
	}

	if request.Year != "" {
		tv.Year = request.Year
	}

	if request.Description != "" {
		tv.Description = request.Description
	}

	response, err := h.TvRepository.UpdateTv(tv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func convertTvResponse(u models.Tv) tvdto.TvResponse {
	return tvdto.TvResponse{
		ID:          u.ID,
		Title:       u.Title,
		Image:       u.Image,
		Year:        u.Year,
		CategoryID:  u.CategoryID,
		Description: u.Description,
		Link:        u.Link,
	}
}
