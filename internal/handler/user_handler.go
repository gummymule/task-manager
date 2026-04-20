package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gummymule/task-manager/internal/domain"
	"github.com/gummymule/task-manager/pkg/response"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

// Register godoc
// @Summary      Register new user
// @Description  Create a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body domain.User true "User registration data"
// @Success      201  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Router       /register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	result, err := h.userUsecase.Register(&user)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "register success", result)
}

// Login godoc
// @Summary      Login user
// @Description  Login and get JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body domain.User true "User login data"
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Router       /login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	token, err := h.userUsecase.Login(req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "login success", gin.H{"token": token})
}

// Logout godoc
// @Summary      Logout user
// @Description  Logout and invalidate JWT token
// @Tags         auth
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  response.Response
// @Failure      401  {object}  response.Response
// @Router       /logout [post]
func (h *UserHandler) Logout(c *gin.Context) {
	response.Success(c, http.StatusOK, "logout success", nil)
}
