package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./names.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type User struct {
	Id           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	HasLoan      bool   `json:"has_loan"`
	HasOtherLoan bool   `json:"has_other_loan"`
}

func GetUsers(count int) ([]User, error) {

	rows, err := DB.Query("SELECT * from users LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		singleUser := User{}
		err = rows.Scan(&singleUser.Id, &singleUser.FirstName, &singleUser.LastName, &singleUser.Email, &singleUser.HasLoan, &singleUser.HasOtherLoan)

		if err != nil {
			return nil, err
		}

		users = append(users, singleUser)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return users, err
}

func GetUserById(id string) (User, error) {

	stmt, err := DB.Prepare("SELECT * from users WHERE id = ?")

	if err != nil {
		return User{}, err
	}

	user := User{}

	sqlErr := stmt.QueryRow(id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.HasLoan, &user.HasOtherLoan)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return User{}, nil
		}
		return User{}, sqlErr
	}
	return user, nil
}

func AddUser(newUser User) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO users (first_name, last_name, email, has_loan, has_other_loan) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newUser.FirstName, newUser.LastName, newUser.Email, newUser.HasLoan, newUser.HasOtherLoan)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func UpdateUser(updatedUser User, id int) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE users SET first_name = ?, last_name = ?, email = ?, has_loan = ?, has_other_loan = ? WHERE Id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(updatedUser.FirstName, updatedUser.LastName, updatedUser.Email, updatedUser.HasLoan, updatedUser.HasOtherLoan, updatedUser.Id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func DeleteUser(userId int) (bool, error) {

	tx, err := DB.Begin()

	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare("DELETE from users where id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
