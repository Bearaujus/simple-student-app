package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/Bearaujus/bthreads/pkg/util"
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

	// Initialize target host
	host := "127.0.0.1:25565"

	// Start service
	util.ClearScreen()
	listener, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Printf("%v %v\n", color.HiRedString("[ ERR ]"), err)
		return
	}
	fmt.Printf("[ %v ] Service started at %v\n", color.HiGreenString(time.Now().Local().Format(time.RFC1123)), color.YellowString(host))
	http.Serve(listener, router)

}
