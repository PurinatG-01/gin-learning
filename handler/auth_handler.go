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
	responder    utils.Responder
}

func NewAuthHandler(loginService service.LoginService, jwtService service.JWTService) *AuthHandler {
	return &AuthHandler{loginService: loginService, jwtService: jwtService, responder: utils.Responder{}}
}

func (s *AuthHandler) Login(c *gin.Context) {
	var credential model.LoginCredentials
	err := c.ShouldBind(&credential)
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	isUserAuthenticated := s.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		s.responder.ResponseSuccess(c, &map[string]interface{}{"token": s.jwtService.GenerateToken(credential.Email, true)})
		return
	} else {
		s.responder.ResponseUnauthorized(c)
		return
	}
}

func (s *AuthHandler) Logout(c *gin.Context) {
	s.responder.ResponseSuccess(c, nil)
	return
}
