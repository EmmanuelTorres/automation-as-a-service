package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DAO interface {
	NewBrandQuery() BrandQuery
	NewCountryQuery() CountryQuery
	NewDesignerQuery() DesignerQuery
	NewProjectQuery() ProjectQuery
	NewUserQuery() UserQuery
}

type dao struct{}

var DB *sql.DB

func pgQb() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(DB)
}

func NewDAO(db *sql.DB) DAO {
	DB = db
	return &dao{}
}

func NewDB() (*sql.DB, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read from a config")
	}
	host := viper.Get("database.host").(string)
	port := viper.Get("database.port").(int)
	user := viper.Get("database.user").(string)
	dbname := viper.Get("database.dbname").(string)
	password := viper.Get("database.password").(string)

	// Starting a database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

func (d *dao) NewBrandQuery() BrandQuery {
	return &brandQuery{}
}

func (d *dao) NewCountryQuery() CountryQuery {
	return &countryQuery{}
}

func (d *dao) NewDesignerQuery() DesignerQuery {
	return &designerQuery{}
}

func (d *dao) NewProjectQuery() ProjectQuery {
	return &projectQuery{}
}

func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{}
}
