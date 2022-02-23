package dto

type Project struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
