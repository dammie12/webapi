package app

import "webapi/controller"

func mapUrls() {
	router.GET("/ping", controller.Ping)
}
