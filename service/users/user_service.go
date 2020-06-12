package users

import (
	"ReshmaKolekar/bookstore_users/domain/users"
	"ReshmaKolekar/bookstore_users/util/errors"
	"log"
)

func GetUser(userId int64) (*users.User, *errors.Rest_Error) {
	result := users.User{ID: userId}
	log.Println("before update:", result)
	if getErr := result.Get(); getErr != nil {
		return nil, getErr
	}
	log.Println("before update:", result)
	return &result, nil
}

func CreateUser(user users.User) (*users.User, *errors.Rest_Error) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	if saveErr := user.Save(); saveErr != nil {
		return nil, saveErr
	}

	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.Rest_Error) {

	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.Status != "" {
			current.Status = user.Status
		}
		if user.Password != "" {
			current.Password = user.Password
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
		current.Password = user.Password
	}

	if updateErr := current.Update(); updateErr != nil {
		return nil, updateErr
	}

	return current, nil
}

func DeleteUser(userID int64) *errors.Rest_Error {
	current, err := GetUser(userID)
	if err != nil {
		return err
	}

	if deleteErr := current.Delete(); deleteErr != nil {
		return deleteErr
	}
	return nil
}

func Search(status string) ([]users.User, *errors.Rest_Error) {
	user := &users.User{}
	return user.Search(status)
}
