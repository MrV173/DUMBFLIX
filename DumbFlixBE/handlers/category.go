package handlers

import (
	categorydto "dumbflix/dto/category"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HandleCategory struct {
	CategoryRepository repository.CategoryRepository
}

func CategoryHandler(categoryRepository repository.CategoryRepository) *HandleCategory {
	return &HandleCategory{categoryRepository}
}

func (h *HandleCategory) FindCategories(c echo.Context) error {
	categories, err := h.CategoryRepository.FindCategories()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: categories,
	})
}

func (h *HandleCategory) GetCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.CategoryRepository.GetCategory(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertCategoryResponse(category)})
}

func (h *HandleCategory) CreateCategory(c echo.Context) error {
	request := new(categorydto.CreateCategoryRequest)
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

	data := models.Category{
		Name: request.Name,
	}

	response, err := h.CategoryRepository.CreateCategory(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func (h *HandleCategory) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CategoryRepository.DeleteCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertCategoryResponse(data),
	})
}

func convertCategoryResponse(u models.Category) categorydto.CategoryResponse {
	return categorydto.CategoryResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
