package main

import (
	"fmt"
	"github.com/grunclepug/networkdiagnostictool/util"
	"github.com/grunclepug/networkdiagnostictool/web"
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

	fmt.Println("[INFO] Starting Request Handler")
	mux.HandleFunc("/", web.RequestHandler)

	fmt.Printf("[INFO] Server listening on http://%v:8080\n", ip)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
}
