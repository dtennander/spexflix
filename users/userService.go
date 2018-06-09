package main

import (
	"database/sql"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"strconv"
	"time"
)

// userDao is a DAO connecting to a table with the form:
//  id | first_name | last_name |           email           |     spex_start      |       creation_date
// ----+------------+-----------+---------------------------+---------------------+----------------------------
//   2 | David      | Tennander | david.tennander@gmail.com | 2015-08-01 00:00:00 | 2018-06-05 13:24:04.19193
//   1 | admin      | admin     | admin@karspexet.se        | 1980-08-01 00:00:00 | 2018-06-05 13:23:58.853459
type userDao struct {
	db *sql.DB
}

type dbConfig struct {
	instanceConnnectionName string
	databaseName            string
	user                    string
	password                string
}

func createUserService(databaseConfig dbConfig) (*userDao, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		databaseConfig.instanceConnnectionName,
		databaseConfig.databaseName,
		databaseConfig.user,
		databaseConfig.password)
	db, err := sql.Open("cloudsqlpostgres", dsn)
	if err != nil {
		return nil, err
	}
	return &userDao{db: db}, nil
}

func (u *userDao) getUser(userId int64) (*User, error) {
	rows, err := u.db.Query("SELECT * FROM users WHERE id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users, err := getUsersFromRows(rows)
	if err != nil {
		return nil, err
	}
	return users[0], nil
}

func getUsersFromRows(rows *sql.Rows) ([]*User, error) {
	var id, firstName, lastName, email, spexStart, creationDate string
	users := make([]*User, 0)
	for rows.Next() {
		rows.Scan(&id, &firstName, &lastName, &email, &spexStart, &creationDate)
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
		users = append(users, user)
	}
	return users, nil
}

func (u *userDao) queryUsers(email string) ([]*User, error) {
	rows, err := u.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users, err := getUsersFromRows(rows)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userDao) postUser(user *User) (int64, error) {
	result, err := u.db.Exec("INSERT INTO (first_name, last_name, email, spex_start) VALUE ($1,$2,$3,$4)",
		user.Name, user.Name, user.Email, user.SpexYears)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}
