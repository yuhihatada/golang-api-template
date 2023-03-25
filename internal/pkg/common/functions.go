package common

import "net/http"

func SetPreflightHeader(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")                                    // origin setting
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")     // method setting
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization") // allow header setting
}
