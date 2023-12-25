package server

import (
	"os"
	"testing"
)

func TestLoadAssets(t *testing.T) {
	const name = "moebooru"
	const write = false
	LoadAssets("../assets/" + name)
	svg := BuildCounterImg("0123456789")
	if !write {
		return
	}
	file, err := os.Create(name + ".svg")
	if err != nil {
		t.Error("Err creating file", err)
	}
	defer file.Close()
	_, err = file.WriteString(svg)
	if err != nil {
		t.Error("Err writing file", err)
	}
}
