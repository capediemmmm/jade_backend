package route

import (
	"jade_backend/api/dto"
	"jade_backend/internal/controller"

	"github.com/gin-gonic/gin"
)

type JadeCtlWrapper struct {
	ctl controller.IJadeController
}

func (w *JadeCtlWrapper) GetAllJade(c *gin.Context) {
	resp, err := w.ctl.GetAllJade(c)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}
