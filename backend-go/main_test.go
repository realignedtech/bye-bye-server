package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHandlerRoutes(t *testing.T) {
	dist := t.TempDir()
	writeFixture(t, dist, "index.html", "<h1>HOME</h1>")
	writeFixture(t, dist, "404.html", "<h1>NOT FOUND</h1>")
	writeFixture(t, dist, "robots.txt", "User-agent: *")
	writeFixture(t, dist, "sitemap.xml", "<urlset></urlset>")
	writeFixture(t, dist, filepath.Join("assets", "app-hash.js"), "console.log('ok')")
	writeFixture(t, dist, filepath.Join("assets", "favicon.png"), "png")

	tests := []struct {
		name   string
		path   string
		status int
		body   string
		cache  string
	}{
		{name: "homepage", path: "/", status: http.StatusOK, body: "HOME", cache: "no-cache"},
		{name: "robots", path: "/robots.txt", status: http.StatusOK, body: "User-agent: *", cache: "no-cache"},
		{name: "sitemap", path: "/sitemap.xml", status: http.StatusOK, body: "<urlset>", cache: "no-cache"},
		{name: "hashed asset", path: "/assets/app-hash.js", status: http.StatusOK, body: "console.log('ok')", cache: "public, max-age=31536000, immutable"},
		{name: "image", path: "/assets/favicon.png", status: http.StatusOK, body: "png", cache: "public, max-age=31536000"},
		{name: "unknown", path: "/not-a-real-page", status: http.StatusNotFound, body: "NOT FOUND", cache: "no-store"},
		{name: "unknown with query", path: "/missing?source=test", status: http.StatusNotFound, body: "NOT FOUND", cache: "no-store"},
	}

	handler := newHandler(dist)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			res := httptest.NewRecorder()
			handler.ServeHTTP(res, req)

			if res.Code != tt.status {
				t.Fatalf("status = %d, want %d; body = %q", res.Code, tt.status, res.Body.String())
			}
			if !strings.Contains(res.Body.String(), tt.body) {
				t.Errorf("body %q does not contain %q", res.Body.String(), tt.body)
			}
			if got := res.Header().Get("Cache-Control"); got != tt.cache {
				t.Errorf("Cache-Control = %q, want %q", got, tt.cache)
			}
		})
	}
}

func writeFixture(t *testing.T, root, name, content string) {
	t.Helper()
	filePath := filepath.Join(root, name)
	if err := os.MkdirAll(filepath.Dir(filePath), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
