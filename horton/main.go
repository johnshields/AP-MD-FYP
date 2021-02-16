/*
 * Horton Golang Server
 * John Shields
 *
 * Main function to start server
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	sw "github.com/GIT_USER_ID/GIT_REPO_ID/go"
	_ "github.com/gin-gonic/gin"
	"log"
)

// CORS are set up in ./go/routers.go
func main() {
	router := sw.NewRouter()
	// log.Println("Horton started on: http://52.51.6.178:8080/api/v1/") // server
	log.Println("Horton started on: http://127.0.0.1:8080/api/v1/") // local
	router.Run()
}
