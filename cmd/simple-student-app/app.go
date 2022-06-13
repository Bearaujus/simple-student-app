package main

import (
	"log"
	"net/http"

	in "github.com/Bearaujus/simple-student-app/init"
	_ "github.com/mattn/go-sqlite3"

	studentHandler "github.com/Bearaujus/simple-student-app/internal/handler/student"
)

func main() {
	// Initialize database
	db, err := in.InitDatabase()
	if err != nil {
		log.Fatalf("Fail to initialize database. %v", err.Error())
		return
	}

	// Initialize handler
	studentHander := studentHandler.NewHandler(db)

	// Initialize router
	router := in.InitRouter(studentHander)
	http.ListenAndServe(":25565", router)
}
