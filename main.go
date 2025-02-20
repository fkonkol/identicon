package main

import (
	"crypto/md5"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
	"strings"

	"github.com/fkonkol/identicon/identicon"
)

func parseIdentifier(r *http.Request) (string, error) {
	identifier := r.PathValue("identifier")

	if !strings.HasSuffix(identifier, ".png") {
		return "", fmt.Errorf("unsupported media type")
	}

	identifier = strings.TrimSuffix(identifier, ".png")

	if strings.TrimSpace(identifier) == "" {
		return "", fmt.Errorf("identifier is required")
	}

	return identifier, nil
}

func getIdenticon(w http.ResponseWriter, r *http.Request) {
	identifier, err := parseIdentifier(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	checksum := md5.Sum([]byte(identifier))
	icon, err := identicon.New(
		identicon.WithSource(checksum[:]),
	)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	bytes, err := icon.Bytes()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(bytes)
}

func main() {
	if err := run(); err != nil {
		trace := string(debug.Stack())
		slog.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

func run() error {
	slog.Info("starting server on :4000")
	http.HandleFunc("/identicons/{identifier}", getIdenticon)
	return http.ListenAndServe(":4000", nil)
}
