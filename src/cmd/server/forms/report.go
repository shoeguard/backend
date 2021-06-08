package forms

type ReportForm struct {
	DeviceInfo string  `example:"00:11:22:33:FF:EE"  valid:"required"`
	Latitude   float64 `example:"37.5428147089301"   valid:"required"`
	Longitude  float64 `example:"126.96631451849314" valid:"required"`
}
