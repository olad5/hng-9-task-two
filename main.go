package main

import (
	"fmt"
	"github.com/olad5/hng-9-task-two/router"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println("Serving on :" + port)
	router := router.Initialize()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
