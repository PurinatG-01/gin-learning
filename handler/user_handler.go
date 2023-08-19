package handler

import (
	"encoding/json"
	model "gin-learning/models"
	"gin-learning/service"
	"gin-learning/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
	jwtService  service.JWTService
	responder   utils.Responder
	paginator   utils.Paginator
}

func NewUserHandler(userService service.UserService, jwtService service.JWTService) *UserHandler {

	return &UserHandler{userService: userService, jwtService: jwtService, responder: utils.Responder{}, paginator: utils.Paginator{}}
}

// UserGetPublic godoc
// @description Get public info user by user id
// @tags User
// @id UserGetPublicHandler
// @produce json
// @param id path int true "User ID"
// @response 200 {object} utils.ApiResponse
// @response 400 {object} utils.ApiResponse
// @Router /user/{id} [get]
func (s *UserHandler) GetPublic(c *gin.Context) {
	param_id := c.Param("id")
	id, param_err := strconv.Atoi(param_id)
	if param_err != nil {
		s.responder.ResponseError(c, param_err.Error())
		return
	}
	user, err := s.userService.GetPublic(id)
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	marshal_event, _ := json.Marshal(&user)
	var data map[string]interface{}
	_ = json.Unmarshal(marshal_event, &data)
	s.responder.ResponseSuccess(c, &data)
	return
}

// UserTickets godoc
// @description Get user's ticket by user id
// @tags User
// @id UserTicketsHandler
// @security JWT
// @produce json
// @param page query int true "page of the list"
// @param limit query int true "limit of the list"
// @response 200 {object} utils.ApiResponse
// @response 400 {object} utils.ApiResponse
// @Router /user/tickets [get]
func (s *UserHandler) Tickets(c *gin.Context) {
	str_user_id := c.GetString("x-user-id")
	user_id, param_err := strconv.Atoi(str_user_id)
	if param_err != nil {
		s.responder.ResponseError(c, param_err.Error())
		return
	}
	paginator_err := s.paginator.Bind(c)
	if paginator_err != nil {
		s.responder.ResponseError(c, paginator_err.Error())
		return
	}
	tickets, tickets_err := s.userService.GetTicketsList(user_id, s.paginator.Page, s.paginator.Limit)
	if tickets_err != nil {
		s.responder.ResponseError(c, tickets_err.Error())
		return
	}
	marshal_event, _ := json.Marshal(&tickets)
	var data map[string]interface{}
	_ = json.Unmarshal(marshal_event, &data)
	s.responder.ResponseSuccess(c, &data)
	return
}

// UserTickets godoc
// @description Get user's transaction by user id
// @tags User
// @id UserTransactionHandler
// @security JWT
// @produce json
// @param page query int true "page of the list"
// @param limit query int true "limit of the list"
// @response 200 {object} utils.ApiResponse
// @response 400 {object} utils.ApiResponse
// @Router /user/tickets [get]
func (s *UserHandler) Transactions(c *gin.Context) {
	str_user_id := c.GetString("x-user-id")
	user_id, param_err := strconv.Atoi(str_user_id)
	if param_err != nil {
		s.responder.ResponseError(c, param_err.Error())
		return
	}
	paginator_err := s.paginator.Bind(c)
	if paginator_err != nil {
		s.responder.ResponseError(c, paginator_err.Error())
		return
	}
	tickets, tickets_err := s.userService.GetTransactionList(user_id, s.paginator.Page, s.paginator.Limit)
	if tickets_err != nil {
		s.responder.ResponseError(c, tickets_err.Error())
		return
	}
	marshal_event, _ := json.Marshal(&tickets)
	var data map[string]interface{}
	_ = json.Unmarshal(marshal_event, &data)
	s.responder.ResponseSuccess(c, &data)
	return
}

// UserTickets godoc
// @description Update user basic info
// @tags User
// @id UserUpdateHandler
// @security JWT
// @produce json
// @param body formData model.UpdateFormUser true "User data to be update"
// @response 204 {object} utils.ApiResponse
// @response 400 {object} utils.ApiResponse
// @Router /user/tickets [get]
func (s *UserHandler) Update(c *gin.Context) {
	str_user_id := c.GetString("x-user-id")
	user_id, param_err := strconv.Atoi(str_user_id)
	if param_err != nil {
		s.responder.ResponseError(c, param_err.Error())
		return
	}
	form_user := model.UpdateFormUser{}
	bind_err := c.ShouldBind(&form_user)
	if bind_err != nil {
		s.responder.ResponseError(c, bind_err.Error())
		return
	}
	update_user, update_err := s.userService.Update(user_id, form_user)
	if update_err != nil {
		s.responder.ResponseError(c, update_err.Error())
		return
	}
	token := s.jwtService.GenerateToken(update_user)
	s.responder.ResponseSuccess(c, &map[string]interface{}{"token": token, "acknowledged": true})
}
