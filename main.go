package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"net/http/cgi"
	"os"
	"time"
)

//go:embed static
var staticFS embed.FS

func main() {
	static, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(os.Stderr, "hello world", time.Now())
	if err := cgi.Serve(http.FileServer(http.FS(static))); err != nil {
		panic(err)
	}
}
