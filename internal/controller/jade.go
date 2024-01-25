package controller

import (
	"jade_backend/api/dto"
	"jade_backend/internal/dao"
	"jade_backend/internal/dao/model"
	"jade_backend/utils/stacktrace"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// >>>>>>>>>>>>>>>>>> Interface  >>>>>>>>>>>>>>>>>>
type IJadeController interface {
	GetAllJade(*gin.Context) (*dto.JadeGetListResp, error)
}

var _ IJadeController = (*JadeController)(nil)

var NewJadeController = func() *JadeController {
	return &JadeController{database: dao.Db}
}

type JadeController struct {
	database *gorm.DB
}

func (c *JadeController) GetAllJade(ctx *gin.Context) (*dto.JadeGetListResp, error) {
	dbQuery := dao.Db.Model(&model.Jade{})

	var total int64
	var Jades []model.Jade
	result := dbQuery.Find(&Jades).Count(&total)

	if result.Error != nil {
		return nil, stacktrace.PropagateWithCode(result.Error, dto.ErrNoShortLink, "Jade not found in database")
	}

	var mappedJades []dto.JadeGetResponse

	for _, jade := range Jades {
		mappedJade := dto.JadeGetResponse{
			Aperture:        jade.Aperture,
			Found:           jade.Found,
			From:            jade.From,
			Graph:           jade.Graph,
			ID:              jade.ID,
			InnerDiameter:   jade.InnerDiameter,
			Material:        jade.Material,
			Number:          jade.Number,
			OutsideDiameter: jade.OutsideDiameter,
			Thick:           jade.Thick,
			Type:            jade.Type,
		}
		mappedJades = append(mappedJades, mappedJade)
	}

	resp := &dto.JadeGetListResp{
		Jades: mappedJades,
	}

	return resp, nil
}
