package server

import (
	"os"
	"testing"
)

func TestLoadAssets(t *testing.T) {
	LoadAssets("../assets/rule34")
	svg := BuildCounterImg("0123456789")
	file, err := os.Create("rule34.svg")
	if err != nil {
		t.Error("Err creating file", err)
	}
	defer file.Close()
	_, err = file.WriteString(svg)
	if err != nil {
		t.Error("Err writing file", err)
	}
}
