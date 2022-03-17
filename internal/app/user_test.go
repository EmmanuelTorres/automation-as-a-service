package app

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/repository"
	"github.com/automation-as-a-service/internal/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var microService *MicroserviceServer

func TestCreateUser(t *testing.T) {
	table := []struct {
		name   string
		person datastruct.Person
		err    error
	}{
		{"Creating admin user", datastruct.Person{FirstName: "Admin", LastName: "User", Email: "emmanueltorres@outlook.com", Password: "password", PhoneNumber: "123456", Role: datastruct.ADMIN}, nil},
		{"Creating regular user", datastruct.Person{FirstName: "Regular", LastName: "User", Email: "emmanueltorres+1@outlook.com", Password: "password", PhoneNumber: "123456", Role: datastruct.USER}, nil},
		{"Creating duplicate user", datastruct.Person{FirstName: "Admin", LastName: "User", Email: "emmanueltorres@outlook.com", Password: "password", PhoneNumber: "123456", Role: datastruct.ADMIN}, errors.New("pq: duplicate key value violates unique constraint \"person_email_key\"")},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			user_id, err := microService.authService.SignUp(tc.person)

			if err != nil {
				if tc.err == nil {
					t.Fatalf("expected user to be created but got error %v", err)
				}

				if tc.err.Error() != err.Error() {
					t.Fatalf("expected error %v but got %v", tc.err, err)
				}

				return
			}

			if user_id == nil {
				t.Fatalf("expected valid user_id but found %v", user_id)
			}

			t.Logf("Created a user with id %d", user_id)
			// w := httptest.NewRecorder()
			// c, _ := gin.CreateTestContext(w)

			// m := NewMicroService()
			// CreateUser(c)
			// req := httptest.NewRequest(http.MethodPost, "localhost:8080/v1/users/", nil)
			// w := httptest.NewRecorder()
			// CreateUser(req, w)
			// res := w.Result()
			// defer res.Body.Close()
			// data, err := ioutil.ReadAll(res.Body)
			// if err != nil {
			// 	t.Fatalf("expected error to not be nil got %v", err)
			// }

			// if data == nil {
			// 	t.Fatalf("expected data not to be nil but got nil")
			// }
		})
	}
}

func TestMain(m *testing.M) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	// Prepare config file
	viper.AddConfigPath("../../")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("cannot read from a  %v", err)
	}

	// Set the database to the testing one
	viper.Set("database.dbname", "aaas_test")

	// Postgres
	db, err := repository.NewDB()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot ping db: %v", err)
	}

	// Set up tables for test execution
	err = createTables(db)
	if err != nil {
		log.Fatalf("could not create tables %v", err)
	}

	// JWT
	signedKeyJwt := viper.Get("jwt.signedKey").(string)
	tokenManager := service.NewTokenManager(signedKeyJwt)

	// Register all services
	dao := repository.NewDAO(db)
	authService := service.NewAuthService(dao, tokenManager)
	brandService := service.NewBrandService(dao)
	countryService := service.NewCountryService(dao)
	designerService := service.NewDesignerService(dao)
	garmentService := service.NewGarmentService(dao)
	projectService := service.NewProjectService(dao)
	userService := service.NewUserService(dao)

	microService = NewMicroService(
		authService,
		brandService,
		countryService,
		designerService,
		garmentService,
		projectService,
		tokenManager,
		userService,
	)

	// Run all tests within this file
	exitVal := m.Run()

	// Tear down tables for test execution
	_ = deleteTables(db)

	os.Exit(exitVal)
}

// Create all the tables on the database for testing purposes
func createTables(db *sql.DB) error {
	_, exists := db.Query("SELECT * FROM person;")
	if exists == nil {
		log.Fatalf("Tried to create a table on a database that already has it")
	}

	_, err := db.Exec(`CREATE TABLE person(
		id bigserial primary key,
		first_name varchar not null,
		last_name varchar not null,
		email varchar not null unique,
		password varchar not null,
		phone_number varchar,
		role varchar not null
	);`)
	if err != nil {
		return err
	}
	return nil
}

func deleteTables(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE person;")
	if err != nil {
		return err
	}

	return nil
}
