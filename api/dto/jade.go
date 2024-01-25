package dto

type JadeGetResponse struct {
	Aperture        float64 `json:"aperture"`           // 孔径
	Found           string  `json:"found"`              // 出土遗址
	From            string  `json:"from"`               // 所属文化
	Graph           string  `json:"graph"`              // 图片
	ID              int64   `json:"id"`                 // ID 编号
	InnerDiameter   float64 `json:"inner_diameter"`     // 内径
	Material        *string `json:"material,omitempty"` // 材质
	Number          string  `json:"number"`             // 器物号
	OutsideDiameter float64 `json:"outside_diameter"`   // 外径
	Thick           float64 `json:"thick"`              // 厚
	Type            string  `json:"type"`               // 类型
}

type JadeGetListResp struct {
	Jades []JadeGetResponse `json:"jades"`
}
