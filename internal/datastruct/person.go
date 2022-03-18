package datastruct

const PersonTableName = "person"

type Person struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     Role   `db:"role"`
}

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)
