package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/picture/", http.StripPrefix("/picture/", http.FileServer(http.Dir("./public/pics"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	//get session cookie
	sessCookie, err := req.Cookie("session")

	//if no session cookie, create
	if err != nil {
		sessCookie = &http.Cookie{
			Name:     "session",
			Value:    uuid.NewV4().String(),
			HttpOnly: true,
		}
		http.SetCookie(w, sessCookie)
	}

	//get previous uploads cookie
	var uploads []string
	uploadsCookie, _ := req.Cookie("uploads")
	if uploadsCookie != nil {
		bytes,_ := base64.StdEncoding.DecodeString(uploadsCookie.Value)
		json.Unmarshal(bytes, &uploads)
	} else {
		uploadsCookie = &http.Cookie{
			Name: "uploads",
		}
	}

	if req.Method == http.MethodPost {
		file, header, err := req.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		extension := getFileExtension(header)
		sha := sha1.New()
		if _, err := io.Copy(sha, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fileName := fmt.Sprintf("%x", sha.Sum(nil)) + "." + extension

		wd, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		newFilePath := filepath.Join(wd, "public", "pics", fileName)
		newFile, err := os.Create(newFilePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer newFile.Close()

		file.Seek(0, 0)
		if _, err := io.Copy(newFile, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		uploads = append(uploads, fileName)
	}

	//set previous uploads cookie
	bytes, err := json.Marshal(&uploads)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	uploadsCookie.Value = base64.StdEncoding.EncodeToString(bytes)
	http.SetCookie(w, uploadsCookie)

	data := struct {
		SessionId      string
		UploadedImages []string
	}{
		sessCookie.Value,
		uploads,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func getFileExtension(header *multipart.FileHeader) string {
	parts := strings.Split(header.Filename, ".")
	return parts[len(parts)-1]
}
