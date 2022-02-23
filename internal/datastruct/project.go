package datastruct

const ProjectTableName = "project"

type Project struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
