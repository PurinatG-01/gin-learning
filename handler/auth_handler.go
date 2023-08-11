package handler

import (
	model "gin-learning/models"
	"gin-learning/service"
	"gin-learning/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	loginService service.LoginService
	jwtService   service.JWTService
	userService  service.UserService
	responder    utils.Responder
}

func NewAuthHandler(loginService service.LoginService, jwtService service.JWTService, userService service.UserService) *AuthHandler {
	return &AuthHandler{loginService: loginService, jwtService: jwtService, userService: userService, responder: utils.Responder{}}
}

// LoginAuth godoc
// @description Login and retreiving JWT token
// @tags Auth
// @id AuthLoginHandler
// @accept mpfd
// @produce json
// @param body formData model.LoginCredentials true "User data for login including `username` and `password`"
// @response 200 {object} utils.ApiResponse
// @response 401 {object} utils.ApiResponse
// @Router /login [post]
func (s *AuthHandler) Login(c *gin.Context) {
	var credential model.LoginCredentials
	bind_err := c.ShouldBind(&credential)
	if bind_err != nil {
		s.responder.ResponseError(c, bind_err.Error())
		return
	}
	user, login_err := s.loginService.LoginUser(credential.Username, credential.Password)
	if login_err != nil {
		s.responder.ResponseUnauthorized(c)
		return
	}
	data := map[string]interface{}{"token": s.jwtService.GenerateToken(user)}
	s.responder.ResponseSuccess(c, &data)
}

func (s *AuthHandler) Logout(c *gin.Context) {
	s.responder.ResponseSuccess(c, nil)
	return
}

// SignupAuth godoc
// @description Sign up for creating user
// @tags Auth
// @id AuthSignupHandler
// @accept mpfd
// @produce json
// @param body formData model.FormUser true "User data to be create"
// @response 200 {object} utils.ApiResponse
// @Router /signup [post]
func (s *AuthHandler) Signup(c *gin.Context) {
	var form_user model.FormUser
	bind_err := c.ShouldBind(&form_user)
	if bind_err != nil {
		s.responder.ResponseError(c, "user formdata binding error")
		return
	}
	result, create_err := s.userService.Create(form_user)
	if create_err != nil {
		s.responder.ResponseError(c, create_err.Error())
		return
	}
	user, login_err := s.loginService.LoginUser(form_user.Username, form_user.Password)
	if login_err != nil {
		s.responder.ResponseUnauthorized(c)
		return
	}
	token := s.jwtService.GenerateToken(user)
	s.responder.ResponseSuccess(c, &map[string]interface{}{"acknowledged": result, "token": token})
}
