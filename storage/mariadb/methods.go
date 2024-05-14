package mariadb

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"isustud.com/m/internal/models"
)

type Storage struct {
	DB *sqlx.DB
}

// type ErrDuplicate struct {
// 	Item string
// }

//	func (e *ErrDuplicate) Error() string {
//		return fmt.Sprintf("%s not found", e.Item)
//	}
var (
	ErrDuplicate = errors.New("connection error")
)

func (s *Storage) GetTag(tag models.Tag) (models.Tag, error) {
	var tagDB models.Tag
	err := s.DB.Get(&tagDB, "SELECT * FROM tags WHERE id=?", tag.ID)
	return tagDB, err
}
func (s *Storage) GetApiKey(token string) (models.Key, error) {
	var key models.Key
	err := s.DB.Get(&key, "SELECT * FROM apikeys WHERE apikey=?", token)
	return key, err
}
func (s *Storage) SetApiKey(login string, token string) error {
	_, err := s.DB.NamedExec(`INSERT INTO apikeys (login,apikey) VALUES (:login,:token)`,
		map[string]interface{}{
			"login": login,
			"token": token,
		})
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		fmt.Println(err.Error())
		if mysqlError.Number == 1062 {
			return fmt.Errorf("failed to add user: %w", ErrDuplicate)
		}
	}
	return err
}

func (s *Storage) SetLogs(login string, alias string) error {
	_, err := s.DB.NamedExec(`INSERT INTO logs (alias,login) VALUES (:login,:alias)`,
		map[string]interface{}{
			"login": login,
			"alias": alias,
		})
	return err
}

// func (s *Storage) GetTag() string {
// 	stmt := `SELECT id,name FROM users`

// 	rows, err := s.Query(stmt)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	usr_list := []*UserAccount{}

// 	for rows.Next() {
// 		usr := &UserAccount{}

// 		err = rows.Scan(&usr.ID, &usr.Username)
// 		if err != nil {
// 			return nil, err
// 		}
// 		usr_list = append(usr_list, usr)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return usr_list, nil
// 	return "pass"
// }
