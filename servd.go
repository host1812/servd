package main

import (
	// "fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	// "time"
)

type Page struct {
	Title    string
	Body     []byte
	HTMLbody template.HTML
}

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/samplepage/", samplePageHandler)
	http.ListenAndServe(":8080", nil)
}

func (p *Page) save() error {
	filename := p.Title // + ".html"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title // + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body, HTMLbody: template.HTML(body)}, nil
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func samplePageHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/samplepage/"):]
	p, err := loadPage("samplepage")
	if err != nil {
		p = &Page{Title: title}
	}

	http.SetCookie(w, &http.Cookie{Name: "sample-auth-cokie", Value: "gsaiscool", Path: "/", Domain: "super.local"})

	renderTemplate(w, "view", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		log.Println("'" + title + "' page not exists")
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	requestCookie, err := r.Cookie("sample-auth-cokie")
	if err != nil {
		log.Println("no cookie")
		http.Redirect(w, r, "/samplepage/", http.StatusFound)
		return
	}
	if requestCookie.Value != "gsaiscool" {
		log.Println("cookie is bad")
	}
	// http.SetCookie(w, &http.Cookie{"test-name", "test-value", "/", "", time.Now().AddDate(0, 0, 1), time.Now().AddDate(0, 0, 1).Format(time.UnixDate), 86400, true, true, "test=tcookie", []string{"test=tcookie"}})
	// http.SetCookie(w, &http.Cookie{"test-name", "test-value", "/", "",)
	// http.SetCookie(w, cookie)
	renderTemplate(w, "view", p)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage("index")
	renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, err := template.ParseFiles(tmpl + ".html")
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// log.Fatal("no '" + tmpl + "' template is detected")
		// return
	}
	// err = t.Execute(w, p)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
