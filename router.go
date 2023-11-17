package main

import (
	"mime"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

func staticHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: try using https://pkg.go.dev/net/http#FileServer to simplify this handler
	path := chi.URLParam(r, "*")
	file, err := nodeModules.ReadFile("node_modules/" + path)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	} else {
		ext := filepath.Ext(path)
		if ext != "" {
			mimeType := mime.TypeByExtension(ext)
			if mimeType != "" {
				w.Header().Add("Content-Type", mimeType)
			}
		}

		w.Header().Add("Cache-Control", "max-age=604800, stale-while-revalidate=86400")
		w.Write(file)
	}
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		hlog.NewHandler(log.Logger),
		hlog.AccessHandler(
			func(r *http.Request, status, size int, duration time.Duration) {
				hlog.FromRequest(r).Info().
					Str("method", r.Method).
					Stringer("url", r.URL).
					Int("status", status).
					Int("size", size).
					Dur("duration", duration).
					Send()
			},
		),
		hlog.RemoteAddrHandler("ip"),
		hlog.UserAgentHandler("user_agent"),
		hlog.RefererHandler("referer"),
		hlog.RequestIDHandler("req_id", "Request-Id"),
		middleware.CleanPath,
	)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := RenderTemplate(w, "items/item.go.html")

		if err != nil {
			log.Logger.Error().Err(err).Send()

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	r.Get("/static/*", staticHandler)

	return r
}
