package route

import (
	"jade_backend/api/dto"
	"jade_backend/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Resp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: "pong~",
	})
}

func SetupJadeController(r *gin.RouterGroup) {
	lcw := JadeCtlWrapper{
		ctl: controller.NewJadeController(),
	}
	p := r.Group("/jade")
	p.GET("/list", lcw.GetAllJade)
}
