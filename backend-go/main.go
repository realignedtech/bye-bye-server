package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	distPath := os.Getenv("DIST_PATH")
	if distPath == "" {
		distPath = "../dist"
	}

	// Check if dist folder exists
	if _, err := os.Stat(distPath); os.IsNotExist(err) {
		log.Fatalf("dist folder not found at %s. Run 'npm run build:frontend' first.", distPath)
	}

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           newHandler(distPath),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("Serving prerendered files from: %s", distPath)
	log.Printf("Server running on: http://localhost:%s", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func newHandler(distPath string) http.Handler {
	fileServer := http.FileServer(http.Dir(distPath))
	index := filepath.Join(distPath, "index.html")
	notFound := filepath.Join(distPath, "404.html")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clean := path.Clean(r.URL.Path)
		if clean == "/" {
			w.Header().Set("Cache-Control", "no-cache")
			http.ServeFile(w, r, index)
			return
		}

		relative := filepath.FromSlash(strings.TrimPrefix(clean, "/"))
		filePath := filepath.Join(distPath, relative)
		if info, err := os.Stat(filePath); err == nil && !info.IsDir() {
			setCacheHeaders(w, clean)
			fileServer.ServeHTTP(w, r)
			return
		}

		if f, err := os.Open(notFound); err == nil {
			defer f.Close()
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Header().Set("Cache-Control", "no-store")
			w.WriteHeader(http.StatusNotFound)
			_, _ = io.Copy(w, f)
			return
		}

		http.Error(w, "404 page not found", http.StatusNotFound)
	})
}

func setCacheHeaders(w http.ResponseWriter, urlPath string) {
	ext := filepath.Ext(urlPath)

	// Hashed assets (JS/CSS with hash in filename)
	if ext == ".js" || ext == ".css" {
		base := filepath.Base(urlPath)
		nameWithoutExt := strings.TrimSuffix(base, ext)
		if strings.Contains(nameWithoutExt, "-") {
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			return
		}
	}

	// Images and fonts
	if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".svg" ||
		ext == ".webp" || ext == ".woff" || ext == ".woff2" || ext == ".ico" {
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		return
	}

	// Everything else
	w.Header().Set("Cache-Control", "no-cache")
}
