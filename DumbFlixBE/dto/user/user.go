package userdto

type CreateUserRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Image    string `json:"image" form:"address" validate:"required"`
	Role     string `json:"role" form:"role"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	FullName string `json:"fullname" form:"fullname"`
	Gender   string `json:"gender" form:"gender"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Image    string `json:"image" form:"image"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Image    string `json:"image"`
	Role     string `json:"role"`
}
