package datastruct

const CountryTableName = "country"

type Country struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
