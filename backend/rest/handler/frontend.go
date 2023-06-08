package handler

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static
var static embed.FS

func ServeFrontend() http.Handler {
	serverRoot, err := fs.Sub(static, "static")
	if err != nil {
		log.Print(err)
		return nil
	}
	var staticFS = http.FS(serverRoot)
	return http.FileServer(staticFS)
}
