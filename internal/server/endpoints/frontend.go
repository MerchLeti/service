package endpoints

import (
	"net/http"
	"os"
	"path"
	"strings"
)

const basePath = "/frontend"

type FrontendEndpoint struct {
	base http.Handler
}

func Frontend() *FrontendEndpoint {
	if !checkExists() {
		return nil
	}
	return &FrontendEndpoint{base: http.FileServer(http.Dir(basePath))}
}

func (f *FrontendEndpoint) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if strings.Contains(path.Base(request.URL.Path), ".") {
		f.base.ServeHTTP(writer, request)
	} else {
		http.ServeFile(writer, request, path.Join(basePath, "index.html"))
	}
}

func checkExists() bool {
	stat, err := os.Stat(basePath)
	return err == nil && stat.IsDir()
}
