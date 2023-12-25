package server

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestAdminServer(t *testing.T) {
	target := "http://localhost:3800/admin/auth"
	payload := []byte(`{"pass", "123456"}`)
	rtype := "application/json"
	r, err := http.Post(target, rtype, bytes.NewBuffer(payload))
	if err != nil {
		t.Error(err)
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Resp Body: %v | %v", body, string(body))

}
