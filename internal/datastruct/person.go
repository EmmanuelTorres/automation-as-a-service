package datastruct

const PersonTableName = "person"

type Person struct {
	ID          int64  `db:"id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Role        Role   `db:"role"`
}

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)
