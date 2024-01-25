package model

const JadeTable = "Jade"

type Jade struct {
	Aperture        float64 `json:"aperture"`                                    // 孔径
	Found           string  `json:"found" gorm:"type:varchar(255)"`              // 出土遗址
	From            string  `json:"from" gorm:"type:varchar(255)"`               // 所属文化
	Graph           string  `json:"graph" gorm:"type:longblob"`                  // 图片
	ID              int64   `json:"id" gorm:"primary_key"`                       // ID 编号
	InnerDiameter   float64 `json:"inner_diameter"`                              // 内径
	Material        *string `json:"material,omitempty" gorm:"type:varchar(255)"` // 材质
	Number          string  `json:"number" gorm:"type:varchar(255)"`             // 器物号
	OutsideDiameter float64 `json:"outside_diameter"`                            // 外径
	Thick           float64 `json:"thick"`                                       // 厚
	Type            string  `json:"type" gorm:"type:varchar(255)"`               // 类型
}
