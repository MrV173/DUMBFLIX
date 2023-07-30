package handlers

import (
	moviedto "dumbflix/dto/movie"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:5000/uploads/"

type HandlerMovie struct {
	MovieRepository repository.MovieRepository
}

func MovieHandler(movieRepository repository.MovieRepository) *HandlerMovie {
	return &HandlerMovie{movieRepository}
}

func (h *HandlerMovie) FindMovies(c echo.Context) error {
	movies, err := h.MovieRepository.FindMovies()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	for i, p := range movies {
		movies[i].Image = path_file + p.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: movies,
	})
}

func (h *HandlerMovie) GetMovie(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	movie, err := h.MovieRepository.GetMovie(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	movie.Image = path_file + movie.Image

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertMovieResponse(movie)})
}

func (h *HandlerMovie) CreateMovie(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	category_id, _ := strconv.Atoi(c.FormValue("category_id"))
	fmt.Println("this is data file", dataFile)

	request := moviedto.CreateMovieRequest{
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

	data := models.Movie{
		Title:       request.Title,
		Image:       request.Image,
		Year:        request.Year,
		CategoryID:  request.CategoryID,
		Description: request.Description,
		Link:        request.Link,
	}

	response, err := h.MovieRepository.CreateMovie(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func (h *HandlerMovie) DeleteMovie(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	movie, err := h.MovieRepository.GetMovie(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.MovieRepository.DeleteMovie(movie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertMovieResponse(data),
	})
}

func (h *HandlerMovie) UpdateMovie(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	movie, err := h.MovieRepository.GetMovie(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	request := new(moviedto.UpdateMovieRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if request.Title != "" {
		movie.Title = request.Title
	}

	if request.Image != "" {
		movie.Image = request.Image
	}

	if request.Year != "" {
		movie.Year = request.Year
	}

	if request.Description != "" {
		movie.Description = request.Description
	}

	if request.Link != "" {
		movie.Link = request.Link
	}

	response, err := h.MovieRepository.UpdateMovie(movie)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func convertMovieResponse(u models.Movie) moviedto.MovieResponse {
	return moviedto.MovieResponse{
		ID:          u.ID,
		Title:       u.Title,
		Image:       u.Image,
		Year:        u.Year,
		CategoryID:  u.CategoryID,
		Description: u.Description,
		Link:        u.Link,
	}
}
