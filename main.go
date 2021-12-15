package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type justFilesFilesystem struct {
    fs http.FileSystem
}

type M map[string]interface{}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
    f, err := fs.fs.Open(name)
    if err != nil {
        return nil, err
    }
    return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
    http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
    return nil, nil
}

func hello(w http.ResponseWriter, req *http.Request) {

    var data = map[string]string{
            "Name":    "john wick",
            "Message": "have a nice day",
    }

    var t, err = template.ParseFiles("views/homepage.html")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

        t.Execute(w, data)
}

func home(w http.ResponseWriter, req *http.Request){
    data := M{"nama" : "Superman"}
    t := template.Must(template.ParseFiles(
        "views/homepage.html",
        "views/footer.html",
        "views/headers.html"))
    err := t.ExecuteTemplate(w, "homepage", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func toko(w http.ResponseWriter, req *http.Request){
    data := M{"nama" : "Toko Emil",
                "alamat" : "Jl. Harapan Bangsa, Labuhan Ratu",
                "spesifikasi" : "terpercaya | Bisa menggandakan uang | Laris manis"}
    t:= template.Must(template.ParseFiles(
        "views/toko.html",
        "views/footer.html",
        "views/headers.html"))
    t.ExecuteTemplate(w, "toko", data)
}

func tokoemil(w http.ResponseWriter, req *http.Request){
    data := M{"nama" : "Toko Emil",
                "alamat" : "Jl. Untung Suropati, Labuhan Ratu",
                "spesifikasi" : "terpercaya | Bisa menggandakan uang | Laris manis",
                "banner" : "/public/media/Toko Emil/banneremil.png"}
    t:= template.Must(template.ParseFiles(
        "views/toko.html",
        "views/footer.html",
        "views/headers.html"))
    t.ExecuteTemplate(w, "toko", data)
}

func indogros(w http.ResponseWriter, req *http.Request){
    data := M{"nama" : "IndoGrosir",
                "alamat" : "Jl. Lintas Sumatra, Lampung",
                "spesifikasi" : "terpercaya | Bisa menggandakan uang | Laris manis",
                "banner" : "/public/media/Indogros/bannerindo.png",
                "logo" : "/public/media/Indogros/logoindo.png"}
    t:= template.Must(template.ParseFiles(
        "views/toko.html",
        "views/footer.html",
        "views/headers.html"))
    t.ExecuteTemplate(w, "toko", data)
}

func admin(w http.ResponseWriter, req *http.Request){
    t := template.Must(template.ParseFiles(
        "views/admin.html",
        "views/headers.html"))
    t.ExecuteTemplate(w, "admin", nil)
}

func main() {
    fs:= justFilesFilesystem{http.Dir("public/")}

    http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(fs)))
    http.HandleFunc("/", home)
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/toko", toko)
    http.HandleFunc("/admin", admin)
    http.HandleFunc("/tokoemil", tokoemil)
    http.HandleFunc("/indogros", indogros)

    //http.Handle("/static/", 
        //http.StripPrefix("/static/", 
            //http.FileServer(http.Dir("public"))))
    fmt.Println("starting web server at http://localhost:8080/")

    http.ListenAndServe(":8080", nil)
}