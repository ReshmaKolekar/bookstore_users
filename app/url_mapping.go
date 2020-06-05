package app

import "ReshmaKolekar/bookstore_users/controller/users"

func mapURLs() {
	router.GET("/user/:userId", users.GetUser)
	router.POST("/user/create", users.CreateUser)
}
