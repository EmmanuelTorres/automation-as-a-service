package datastruct

const BrandTableName = "brand"

type Brand struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	CountryID int64  `db:"country_id"`
}
