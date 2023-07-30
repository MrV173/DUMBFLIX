package handlers

import (
	dto "dumbflix/dto/result"
	userdto "dumbflix/dto/user"
	"dumbflix/models"
	"dumbflix/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var profile_file = "http://localhost:5000/uploads/"

type HandlerUser struct {
	UserRepository repository.UserRepository
}

func UserHandler(userRepository repository.UserRepository) *HandlerUser {
	return &HandlerUser{userRepository}
}

func (h *HandlerUser) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: users,
	})
}

func (h *HandlerUser) GetUser(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userid := userLogin.(jwt.MapClaims)["id"].(float64)
	user := int(userid)

	User, err := h.UserRepository.GetUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	User.Image = profile_file + User.Image

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponse(User)})
}

func (h *HandlerUser) CreateUser(c echo.Context) error {
	request := new(userdto.CreateUserRequest)
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

	data := models.User{
		FullName: request.FullName,
		Email:    request.Email,
		Password: request.Password,
		Gender:   request.Gender,
		Phone:    request.Phone,
		Address:  request.Address,
		Image:    request.Image,
		Role:     "customer",
	}

	response, err := h.UserRepository.CreateUser(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})
}

func (h *HandlerUser) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponse(data),
	})
}

func (h *HandlerUser) UpdateUser(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userid := userLogin.(jwt.MapClaims)["id"].(float64)
	user := int(userid)

	User, err := h.UserRepository.GetUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	dataFile := c.Get("dataFile").(string)

	request := userdto.UpdateUserRequest{
		Image: dataFile,
	}

	if request.Image != "" {
		User.Image = request.Image
	}

	response, err := h.UserRepository.UpdateUser(User)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})

}

func convertResponse(u models.User) userdto.UserResponse {
	return userdto.UserResponse{
		ID:       u.ID,
		Email:    u.Email,
		FullName: u.FullName,
		Gender:   u.Gender,
		Address:  u.Address,
		Phone:    u.Phone,
		Image:    u.Image,
	}
}
