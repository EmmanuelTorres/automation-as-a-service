package datastruct

const CountryTableName = "country"

type Country struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
