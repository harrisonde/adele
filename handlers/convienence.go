package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"git.int.86labs.cloud/harrisonde/adele-framework"
)

func (h *Handlers) render(w http.ResponseWriter, r *http.Request, template string, variables, data interface{}) error {
	return h.App.Render.Page(w, r, template, variables, data)
}

func (h *Handlers) renderInertia(w http.ResponseWriter, r *http.Request, template string) error {
	return h.App.Render.InertiaPage(w, r, template)
}

func (h *Handlers) sessionPut(ctx context.Context, key string, val interface{}) {
	h.App.Session.Put(ctx, key, val)
}

func (h *Handlers) sessionHas(ctx context.Context, key string) bool {
	return h.App.Session.Exists(ctx, key)
}

func (h *Handlers) sessionGet(ctx context.Context, key string) interface{} {
	return h.App.Session.Get(ctx, key)
}

func (h *Handlers) sessionRemove(ctx context.Context, key string) {
	h.App.Session.Remove(ctx, key)
}

func (h *Handlers) sessionRenew(ctx context.Context) error {
	return h.App.Session.RenewToken(ctx)
}

func (h *Handlers) sessionDestroy(ctx context.Context) error {
	return h.App.Session.Destroy(ctx)
}

func (h *Handlers) randomString(n int) string {
	return h.App.RandomString(n)
}

func (h *Handlers) isAlpha(s string) bool {
	const alpha = "abcdefghijklmnopqrstuvwxyz"
	for _, char := range s {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

func (h *Handlers) isAlphaAnd(s, r string) bool {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	alpha = alpha + string(r)
	for _, char := range s {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

func (h *Handlers) encrypt(text string) (string, error) {
	enc := adele.Encryption{Key: []byte(h.App.EncryptionKey)}

	encrypted, err := enc.Encrypt(text)
	if err != nil {
		return "", err
	}
	return encrypted, nil
}

func (h *Handlers) decrypt(crypto string) (string, error) {
	enc := adele.Encryption{Key: []byte(h.App.EncryptionKey)}

	decrypted, err := enc.Decrypt(crypto)
	if err != nil {
		return "", err
	}
	return decrypted, nil
}

func (h *Handlers) HealthStatus(w http.ResponseWriter, r *http.Request) {
	type PingStatus struct {
		Time        time.Time `json:"time"`
		Description string    `json:"description"`
	}

	res := PingStatus{
		Time:        time.Now(),
		Description: "server is running",
	}

	w.Header().Set("Content-Type", "application/json")

	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Write(jsonRes)
}
