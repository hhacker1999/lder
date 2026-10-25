package router

import (
	"io/fs"
	"lder/web"
	"net/http"
	"strings"
)

func (r *HttpRouter) mountStatic() {
	dist, err := fs.Sub(web.Dist, "dist")
	if err != nil {
		panic("embed web/dist: " + err.Error())
	}
	fileServer := http.FileServer(http.FS(dist))

	r.router.Get("/*", func(w http.ResponseWriter, req *http.Request) {
		path := strings.TrimPrefix(req.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}

		if _, err := dist.Open(path); err != nil {
			data, err := fs.ReadFile(dist, "index.html")
			if err != nil {
				http.NotFound(w, req)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(data)
			return
		}

		req.URL.Path = "/" + path
		fileServer.ServeHTTP(w, req)
	})
}
