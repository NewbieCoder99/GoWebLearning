package main

import (
    "fmt"
    "html/template"
    "net/http"
)

type PageVariables struct {
	Nama string
	About string
}

func main() {
    http.HandleFunc("/", HomePage)
    http.ListenAndServe("127.0.0.1:8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
    
    switch r.Method {
        case "GET" :
            // Mengambil variabel dari struct
            // HomePageVars := PageVariables {
            //     Nama: "Qudrat Nurfajar Yasin Sutisna",
            //     About: "Programmer & Music Creator",
            // }
            t, err := template.ParseFiles("resources/view/home.html")
            // err = t.Execute(w, HomePageVars)
            err = t.Execute(w, nil)
            if err != nil { // if there is an error
                fmt.Fprintf(w, "Hello World")
            }
        case "POST" :
            err := r.ParseForm();
            if err != nil {
                fmt.Fprintf(w, "ParseForm() err: %v", err)
                return
            }
            fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
            email := r.FormValue("email")
            password := r.FormValue("password")
            fmt.Fprintf(w, "Email = %s\n", email)
            fmt.Fprintf(w, "Password = %s\n", password)
        default :
    }
}
