// Code generated by hertz generator. DO NOT EDIT.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	router "hertz-examples/hz_demo/hertz-server/biz/router"
)

// register registers all routers.
func register(r *server.Hertz) {

	// Routing defined in idl
	router.GeneratedRegister(r)

	// User-defined routing
	customizedRegister(r)
}