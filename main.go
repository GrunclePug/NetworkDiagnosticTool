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
	ips := util.LocalIPs()
	ip = ips[0]
	if len(ips) > 1 {
		fmt.Println("[WARN] Multiple IP's detected, this may result in the incorrect IP being listed for the web server, listing all IP's: ", ips)
	}

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
