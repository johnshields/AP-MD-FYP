/*
 * Horton Golang Server
 * John Shields
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */
package main

import (
	sw "github.com/GIT_USER_ID/GIT_REPO_ID/go"
	"log"
)

// Main function to start server
// CORS are set up in ./go/routers.go
func main() {
	log.Printf("Server started")
	log.Println("Server started on: http://52.51.6.178:8080")
	log.Println("Server started on: http://localhost:8080")
	router := sw.NewRouter()
	router.Run()
}
