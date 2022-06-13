package main

import (
	"fmt"
	"log"
	"net/http"

	in "github.com/Bearaujus/simple-student-app/init"
	"github.com/fatih/color"
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
	fmt.Println(color.HiYellowString("Services started at 'localhost:25565'"))
	http.ListenAndServe(":25565", router)
}
