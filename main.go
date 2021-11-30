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
    t := template.Must(template.ParseFiles("views/homepage.html"))
    t.Execute(w, nil)
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func toko(w http.ResponseWriter, req *http.Request){
    t:= template.Must(template.ParseFiles("views/toko.html"))
    t.Execute(w, nil)
}

func main() {
    fs:= justFilesFilesystem{http.Dir("public/")}

    http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(fs)))
    http.HandleFunc("/", home)
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/toko", toko)

    //http.Handle("/static/", 
        //http.StripPrefix("/static/", 
            //http.FileServer(http.Dir("public"))))
    fmt.Println("starting web server at http://localhost:8080/")

    http.ListenAndServe(":8080", nil)
}