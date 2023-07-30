package handlers

import (
	dto "dumbflix/dto/result"
	tvepisodedto "dumbflix/dto/tv-episode"
	"dumbflix/models"
	"dumbflix/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var eps_image = "http://localhost:5000/uploads"

type HandlerEps struct {
	EpsRepository repository.EpisodeTv
}

func EpsHandler(epsRepository repository.EpisodeTv) *HandlerEps {
	return &HandlerEps{epsRepository}
}

func (h *HandlerEps) FindEps(c echo.Context) error {
	epss, err := h.EpsRepository.FindEps()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	for i, p := range epss {
		epss[i].Image = eps_image + p.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: epss,
	})
}

func (h *HandlerEps) GetEps(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	eps, err := h.EpsRepository.GetEps(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	eps.Image = eps_image + eps.Image

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertEpsResponse(eps)})
}

func (h *HandlerEps) CreateEps(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	tv_id, _ := strconv.Atoi(c.FormValue("tv_id"))

	request := tvepisodedto.CreateEpisode{
		Title: c.FormValue("title"),
		Image: dataFile,
		TvID:  tv_id,
		Link:  c.FormValue("link"),
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	data := models.EpisodeTV{
		Title: request.Title,
		Image: request.Image,
		TvID:  request.TvID,
		Link:  request.Link,
	}

	response, err := h.EpsRepository.CreateEps(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func convertEpsResponse(u models.EpisodeTV) tvepisodedto.EpisodeTVResponse {
	return tvepisodedto.EpisodeTVResponse{
		Title: u.Title,
		Image: u.Image,
		TvID:  u.TvID,
		Link:  u.Link,
	}
}
