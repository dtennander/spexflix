package authentication

import (
	"database/sql"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
)

// authRecord represents a Password bound to a User.
type authRecord struct {
	id int64
	userId int64
	hash string
	salt string
}

// authDao lets the User store and retiree authRecords in the database.
// The table structure is on the format:
//
//      auth_id   |   user_id |     hash     |     salt     | creation_date
//   -------------+-----------+--------------+--------------+---------------
//     bigserial  |   bigint  | varchar(255) | varchar(255) |  time_stamp
//
type authDao struct {
	db *sql.DB
}

func (dao *authDao) GetHash(i int64) ([]byte, error) {
	rows, err := dao.db.Query("SELECT hash FROM auth WHERE user_id = $1", i)
	if err != nil {
		return nil, err
	}
	var hash string
	for rows.Next()  {
		err := rows.Scan(&hash)
		if err != nil {
			return nil, err
		}
	}
	return []byte(hash), nil
}

// DbConfig is the configuration passed to createAuthDao to configure the Dao.
type DbConfig struct {
	InstanceConnnectionName string
	User                    string
	Password                string
}

func CreateAuthDao(databaseConfig DbConfig) (*authDao, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		databaseConfig.InstanceConnnectionName,
		"auth",
		databaseConfig.User,
		databaseConfig.Password)
	db, err := sql.Open("cloudsqlpostgres", dsn)
	if err != nil {
		return nil, err
	}
	return &authDao{db: db}, nil
}
