package main

import (
	"fmt"
	"github.com/grunclepug/basicgowebapp/util"
	"github.com/grunclepug/basicgowebapp/web"
	"log"
	"net"
	"net/http"
)

const port string = "8080"

var ip net.IP

func main() {
	ip = util.GetLocalIPs()[0]

	fmt.Printf("[INFO] Starting server on %v:8080\n", ip)
	mux := http.NewServeMux()

	fmt.Println("[INFO] Serving static content")
	mux.Handle("/web/static/", http.StripPrefix("/web/static/", http.FileServer(http.Dir("web/static"))))

	// Consider having multiple handlers for scalability. In this simple example it's easier to handle it all in one handler
	// mux.HandleFunc("/", web.WebHandler)
	// mux.HandleFunc("/send", web.UserHandler)
	fmt.Println("[INFO] Starting User Handler")
	mux.HandleFunc("/", web.UserHandler)

	fmt.Printf("[INFO] Server listening on http://%v:8080\n", ip)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
}
