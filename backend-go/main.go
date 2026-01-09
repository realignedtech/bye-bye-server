package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

	log.Printf("Serving static files from: %s", distPath)
	log.Printf("Server running on: http://localhost:%s", port)

	http.HandleFunc("/", spaHandler(distPath))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// spaHandler serves static files with SPA fallback
func spaHandler(distPath string) http.HandlerFunc {
	fileServer := http.FileServer(http.Dir(distPath))

	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(distPath, r.URL.Path)

		// Check if file exists
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			// Serve index.html for SPA routing
			w.Header().Set("Cache-Control", "no-cache")
			http.ServeFile(w, r, filepath.Join(distPath, "index.html"))
			return
		}

		// Set cache headers for static assets
		setCacheHeaders(w, r.URL.Path)
		fileServer.ServeHTTP(w, r)
	}
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
