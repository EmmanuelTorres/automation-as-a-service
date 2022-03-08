package datastruct

const GarmentTableName = "garment"

type Garment struct {
	ID         int64  `json:"id"`
	Code       string `json:"code"`
	DesignerID int64  `json:"designer_id"`
	BrandID    int64  `json:"brand_id"`
}
