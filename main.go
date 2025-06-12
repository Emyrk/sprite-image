package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	http.HandleFunc("/generate", generateHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	body := r.URL.Query().Get("body")
	head := r.URL.Query().Get("head")
	legs := r.URL.Query().Get("legs")

	if body == "" || head == "" || legs == "" {
		http.Error(w, "Missing one or more query parameters: body, head, legs", http.StatusBadRequest)
		return
	}

	tmpDir := os.TempDir()
	outFile := filepath.Join(tmpDir, "out.png")

	cmd := exec.Command(
		"~/.cargo/bin/lpcg-build",
		"./spritesheets",
		fmt.Sprintf("body::bodies::%s", body),
		fmt.Sprintf("head::heads::%s", head),
		fmt.Sprintf("legs::pantaloons::%s", legs),
		outFile,
	)

	// Use shell to expand the ~ (home dir)
	cmd.Env = append(os.Environ(), "HOME="+os.Getenv("HOME"))
	cmd.Path = "/bin/bash"
	cmd.Args = []string{"bash", "-c", cmd.String()}

	log.Printf("Executing command: %s", cmd.String())

	err := cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Command failed: %v", err), http.StatusInternalServerError)
		return
	}
	defer os.Remove(outFile)

	file, err := os.Open(outFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to open output file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "attachment; filename=out.png")
	_, err = io.Copy(w, file)
	if err != nil {
		log.Printf("Failed to send file: %v", err)
	}
}
