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
	ModifyJade(*gin.Context, *dto.JadeModifyReq) (error)
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
		return nil, stacktrace.PropagateWithCode(result.Error, dto.ErrNoJade, "Jade not found in database")
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

func (c *JadeController) ModifyJade(ctx *gin.Context, req *dto.JadeModifyReq) (error) {
	dbQuery := dao.Db.Model(&model.Jade{})
	NewJade := model.Jade{
		Aperture:        req.Aperture,
		Found:           req.Found,
		From:            req.From,
		Graph:           req.Graph,
		ID:              req.ID,
		InnerDiameter:   req.InnerDiameter,
		Material:        req.Material,
		Number:          req.Number,
		OutsideDiameter: req.OutsideDiameter,
		Thick:           req.Thick,
		Type:            req.Type,
	}
	var jade model.Jade
	result := dbQuery.Where(&model.Jade{ID: req.ID}).First(&jade).Updates(&NewJade)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return stacktrace.PropagateWithCode(result.Error, dto.ErrJadeNotFound, "The jade to be modified is not found in database")
		}
		return stacktrace.PropagateWithCode(result.Error, dto.InternalError, "Failed to modify jade information in database")
	}

	return nil
}