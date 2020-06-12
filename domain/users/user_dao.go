package users

import (
	"ReshmaKolekar/bookstore_users/datasource/mysql/users_db"
	"ReshmaKolekar/bookstore_users/util/errors"
	"ReshmaKolekar/bookstore_users/util/errors/date_utils"
	"ReshmaKolekar/bookstore_users/util/mysql_utils"
	"fmt"
	"log"
)

// var (
// 	userDB = make(map[int64]*User)
// )
const (
	queryInsertUser   = "INSERT INTO users(firstName,lastName,email,created,status,password)values(?,?,?,?,?,?)"
	queryGetUser      = "SELECT *FROM users where id =?"
	queryUpdateUser   = "UPDATE users SET firstName=?,lastName=?,email=?,status=?,password=? where id=?"
	queryDeleteUser   = "DELETE FROM users where id=?"
	queryFindByStatus = "SELECT * from users where status=?"
	//errorNoRows      = "no rows in result set"
	//indexUniqueEmail = "email_UNIQUE"
)

func (user *User) Get() *errors.Rest_Error {
	// if result := userDB[user.ID]; result != nil {
	// 	user.ID = result.ID
	// 	user.FirstName = result.FirstName
	// 	user.LastName = result.LastName
	// 	user.Email = result.Email
	// 	user.DateCreated = result.DateCreated

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	getResult := stmt.QueryRow(user.ID)
	if err := getResult.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
		return mysql_utils.ParseError(err)
		// if strings.Contains(err.Error(), errorNoRows) {
		// 	return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
		// }
		// //return errors.NewNotFoundError(fmt.Sprintf("not found %s", err.Error()))
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d", user.ID))
	}
	return nil
}

func (user *User) Save() *errors.Rest_Error {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	// current := userDB[user.ID]
	// if current != nil {
	// 	if current.Email == user.Email {
	// 		return errors.NewNotFoundError(fmt.Sprintf("Email %s is already registered", user.Email))
	// 	}
	// 	return errors.NewNotFoundError(fmt.Sprintf("user %d already exist", user.ID))
	// }
	// user.DateCreated = date_utils.GetNowString()
	// userDB[user.ID] = user
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		return mysql_utils.ParseError(err)
		// 	if strings.Contains(err.Error(), indexUniqueEmail) {
		// 		return errors.NewBadRequestError(fmt.Sprintf("Email %s is already exist", user.Email))
		// 	}
		// 	return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user %s", err.Error()))
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
		//return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user %s", err.Error()))
	}
	user.ID = userID
	return nil
}

func (user *User) Update() *errors.Rest_Error {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	if _, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.ID); err != nil {
		return mysql_utils.ParseError(err)
		// 	if strings.Contains(err.Error(), indexUniqueEmail) {
		// 		return errors.NewBadRequestError(fmt.Sprintf("Email %s is already exist", user.Email))
		// 	}
		// 	return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user %s", err.Error()))
	}
	return nil
}

func (user *User) Delete() *errors.Rest_Error {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	if _, err := stmt.Exec(user.ID); err != nil {
		return mysql_utils.ParseError(err)
		// 	if strings.Contains(err.Error(), indexUniqueEmail) {
		// 		return errors.NewBadRequestError(fmt.Sprintf("Email %s is already exist", user.Email))
		// 	}
		// 	return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user %s", err.Error()))
	}
	return nil
}

func (user *User) Search(status string) ([]User, *errors.Rest_Error) {

	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	rows, err := stmt.Query(status)
	log.Println("serachquery: ", queryFindByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, user)
	}

	if len(result) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching with status %s", status))
	}
	return result, nil
}
