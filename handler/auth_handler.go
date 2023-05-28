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
	s.responder.ResponseSuccess(c, &map[string]interface{}{"acknowledged": result})
}
