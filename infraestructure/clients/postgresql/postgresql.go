package postgresql

import (
	"fmt"

	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "sBLRWyyPsInwHftmHAWmYJURGWBGFpLs"
	password = "tuXL3XSF8O7tsGrcGHoMos4tVNtL3tnrRshSCZokGnIfk4ArDyzaa297k2WgQPSL"
	dbname   = "admin"
)

func StartGorm() *gorm.DB {
	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(pgdriver.Open(psqlConnectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
