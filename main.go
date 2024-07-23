package main

import (
	"fmt"
	"os"
	"sync"
	"task-management-system/internal/database"
	"task-management-system/internal/server"
)

const port = 8080

func main() {
	myServer := server.NewHttpServer()
	myServer.SetPort(port)

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Error connecting to database: %w", err)
		os.Exit(1)
	}
	mux := server.CreateRoutes(myServer, db)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := myServer.StartServer(mux)
		if err != nil {
			fmt.Println("Error starting server: %w", err)
			os.Exit(1)
		}
	}()
	fmt.Printf("Server started on port %d\n", port)
	wg.Wait()
}
