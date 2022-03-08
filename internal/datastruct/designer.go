package datastruct

const DesignerTableName = "designer"

type Designer struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	CountryID int64  `db:"country_id"`
}
