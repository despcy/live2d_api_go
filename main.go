package main

import (
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	port := ":" + os.Args[1]
	https := os.Args[2]
	fs := http.FileServer(http.Dir("./model"))

	mux := http.NewServeMux()
	mux.Handle("/live2Dmodel/", http.StripPrefix("/live2Dmodel/", fs))
	mux.HandleFunc("/get/", GetModel)
	mux.HandleFunc("/switch/", SwitchModel)
	mux.HandleFunc("/rand_textures/", RandomTextures)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		AllowOriginRequestFunc: func(r *http.Request, origin string) bool { return true },
		Debug:                  false,
	})
	handler := c.Handler(mux) //add CORS support
	if https == "https" {
		http.ListenAndServeTLS(port, "server.crt", "server.key", handler)
	} else {
		http.ListenAndServe(port, handler)
	}
}
