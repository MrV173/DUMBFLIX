package authdto

type AuthRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Image    string `json:"image" form:"image"`
	Role     string `json:"role" form:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
