package app

import "ReshmaKolekar/bookstore_users/controller/users"

func mapURLs() {
	router.GET("/user/:userID", users.Get)
	router.POST("/user/create", users.Create)
	router.PUT("/user/:userID", users.Update)
	router.PATCH("/user/:userID", users.Update)
	router.DELETE("/user/:userID", users.Delete)
	router.GET("/internal/user/serach", users.Search)

}
