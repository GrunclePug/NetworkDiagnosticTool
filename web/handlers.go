package web

import (
	"encoding/json"
	"fmt"
	"github.com/grunclepug/networkdiagnostictool/util"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	SysInfo util.SysInfo
	NetInfo util.NetInfo
}

// RequestHandler Handles HTTP Requests
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] New Request, Target: %v Source: %v\n", r.URL.Path, r.RemoteAddr)

	// Build Data Struct
	sysInfo := util.GetSysInfo()
	netInfo := util.GetNetInfo()
	data := Data{
		SysInfo: sysInfo,
		NetInfo: netInfo,
	}

	// Determine if request is for Dashboard or JSON response
	switch r.URL.Path {
	case "/": // Dashboard
		fmt.Println("[INFO] Serving Dashboard")
		templating(w, "web/static/index.html", data)
	case "/json": // JSON API
		if r.Method == "GET" {
			// Marshal Data Struct to JSON
			jsonData, err := json.Marshal(data)
			if err != nil {
				marshalError := "[ERROR] Error Generating JSON Response"
				fmt.Println(marshalError)
				http.Error(w, marshalError, http.StatusInternalServerError)
				return
			}

			// Serve JSON
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		} else { // Fail on POST
			methodError := "[ERROR] Incorrect HTTP Request Method Error"
			fmt.Println(methodError)
			http.Error(w, methodError, http.StatusInternalServerError)
			return
		}
	}
}

// templating HTML Serving using Templates
func templating(w http.ResponseWriter, fileName string, data interface{}) {
	// Read HTML file
	t := template.Must(template.ParseFiles(fileName))

	// Serve HTML data
	err := t.Execute(w, data)
	if err != nil {
		fmt.Println("[ERROR] HTML Template Execution Error: ", err.Error())
		log.Fatal(err)
	}
}
