package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

type AsciiArt struct {
	TextArea string
	Ascii    string
	Output   []byte
}

var (
	t *template.Template
	s AsciiArt
)

func main() {
	t, _ = template.ParseGlob("templates/*.html")

	fs := http.FileServer(http.Dir("./text/"))
	http.Handle("/text/", http.StripPrefix("/text/", fs))
	http.HandleFunc("/", renderTemplate)
	http.HandleFunc("/ascii-art", processposthandler)
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("ascii.html")
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return

	}
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	err = t.Execute(w, s)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return

	}
}

func processposthandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s.TextArea = r.FormValue("textarea")

	s.Ascii = r.FormValue("ascii")

	if r.FormValue("textarea") == "" {
		http.Error(w, "400 bad request\nrequested character(s) are not in the valid range", http.StatusBadRequest)
		return
	}

	switch s.Ascii {
	case "standard":
	case "shadow":
	case "thinkertoy":
	default:
		s.Ascii = "standard"
	}

	var style string = s.Ascii

	if style != "standard" && style != "shadow" && style != "thinkertoy" {

		http.Error(w, "404 the banner file has not been found or you did not send any data\ntry to start from localhost:8080", http.StatusNotFound)
		return
	}

	data, _ := Banner(style)
	output := s.TextArea
	bannerMap := Array(data)

	ascii_Output := Print(output, bannerMap)
	defer data.Close()

	s.Output = []byte(ascii_Output)

	t, err := template.ParseFiles("ascii.html")
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return

	}

	err = t.Execute(w, s)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
}

func Banner(style string) (file *os.File, err error) {
	b, err := os.Open(fmt.Sprintf("text/%s.txt", style))
	if err != nil {
		panic("404 the banner file has not been found or you did not send any data\ntry to start from localhost:8080")
	}
	return b, err
}

func Array(file *os.File) map[rune][]string {
	bannerMap := make(map[rune][]string)
	var lines []string
	scanner := bufio.NewScanner(io.Reader(file))
	firstLine := false
	char := ' '
	for scanner.Scan() {
		if scanner.Text() == "" && firstLine {
			bannerMap[char] = lines
			lines = nil
			char++
		} else if firstLine {
			lines = append(lines, scanner.Text())
		}
		firstLine = true
	}
	bannerMap[char] = lines
	return bannerMap
}

func Print(str string, banner map[rune][]string) string {
	output := ""
	list := Split(str)
	for _, word := range list {
		if word == "" {
			output += string(rune(10))
		} else {
			for i := 0; i < 8; i++ {
				line := ""
				for _, r := range word {
					line += banner[r][i]
				}
				output += line + string(rune(10))
			}
		}
	}

	return output

}
func Split(str string) []string {
	answer := ""

	b := strings.Split(str, "\r\n")

	for _, j := range b {
		if j == "" {
			answer += string(rune(10))
		}
	}
	return b
}
