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

func (s *UtilityHandler) Shuffle(c *gin.Context) {
	str := c.Query("list")
	rand.Seed(time.Now().Unix())
	list := strings.Split(str, ",")
	data := map[string]interface{}{"result": fmt.Sprint(list[rand.Intn(len(list))])}
	s.responder.ResponseSuccess(c, &data)
	return
}
