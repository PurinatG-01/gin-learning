package handler

import (
	"fmt"
	"gin-learning/utils"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UtilityHandler struct {
	responder utils.Responder
}

func NewUtilityHandler() *UtilityHandler {
	return &UtilityHandler{responder: utils.Responder{}}
}

// UtiltyRandom godoc
// @description Random items in accepting list
// @tags Utility
// @id UtiltyRandomHandler
// @produce json
// @param list query string true "list of items to random in comma separated string formatted, example `1,2.3,4,5` "
// @response 200 {object} utils.ApiResponse
// @Router /utility/random [get]
func (s *UtilityHandler) Random(c *gin.Context) {
	str := c.Query("list")
	rand.Seed(time.Now().Unix())
	list := strings.Split(str, ",")
	data := map[string]interface{}{"result": fmt.Sprint(list[rand.Intn(len(list))])}
	s.responder.ResponseSuccess(c, &data)
	return
}
