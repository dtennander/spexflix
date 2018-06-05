package main

import (
	"database/sql"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"strconv"
	"time"
)

type userService struct {
	db *sql.DB
}

type dbConfig struct {
	instanceConnnectionName string
	databaseName            string
	user                    string
	password                string
}

func createUserService(databaseConfig dbConfig) (*userService, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		databaseConfig.instanceConnnectionName,
		databaseConfig.databaseName,
		databaseConfig.user,
		databaseConfig.password)
	db, err := sql.Open("cloudsqlpostgres", dsn)
	if err != nil {
		return nil, err
	}
	return &userService{db: db}, nil
}

func (u *userService) getUser(userId int64) (*User, error) {
	rows, err := u.db.Query("SELECT * FROM users WHERE id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var id, firstName, lastName, email, spexStart, creationDate string
	for rows.Next() {
		rows.Scan(&id, &firstName, &lastName, &email, &spexStart, &creationDate)
	}
	spexTime, err := time.Parse(time.RFC3339, spexStart)
	if err != nil {
		return nil, err
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	user := &User{
		Id:        idInt,
		Email:     email,
		Name:      firstName,
		SpexYears: int(time.Now().Sub(spexTime).Hours() / (24 * 365)),
	}
	return user, nil
}
