package main

import "net/http"

func secureHeaders(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		h := w.Header()

		h.Set("Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		h.Set("Referrer-Policy", "origin-when-cross-origin")
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("X-Frame-Options", "deny")
		h.Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}
