package users

import (
	domainUser "ReshmaKolekar/bookstore_users/domain/users"
	services "ReshmaKolekar/bookstore_users/service/users"
	"ReshmaKolekar/bookstore_users/util/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserID(userIDParam string) (int64, *errors.Rest_Error) {
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("userId should be a number")
	}
	return userID, nil
}

func Create(c *gin.Context) {
	var user domainUser.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid JSON")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)

	//c.String(http.StatusNotImplemented, "implement me!")
}

func Get(c *gin.Context) {
	//retrive the userid and convert it from string to int
	userID, userIDErr := getUserID(c.Param("userID"))
	if userIDErr != nil {
		c.JSON(userIDErr.Status, userIDErr)
		return
	}

	result, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Update(c *gin.Context) {
	var user domainUser.User

	userID, userIDErr := getUserID(c.Param("userID"))
	if userIDErr != nil {
		c.JSON(userIDErr.Status, userIDErr)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid JSON")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID
	isPartial := c.Request.Method == http.MethodPatch
	result, saveErr := services.UpdateUser(isPartial, user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)

	//c.String(http.StatusNotImplemented, "implement me!")
}
func Delete(c *gin.Context) {
	//var user domainUser.User

	userID, userIDErr := getUserID(c.Param("userID"))
	if userIDErr != nil {
		c.JSON(userIDErr.Status, userIDErr)
		return
	}
	if delErr := services.DeleteUser(userID); delErr != nil {
		c.JSON(delErr.Status, delErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	result, searchErr := services.Search(status)
	if searchErr != nil {
		c.JSON(searchErr.Status, searchErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
