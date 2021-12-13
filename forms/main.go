package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    /*"encoding/json"
    "golang.org/x/crypto/bcrypt"
    "os"*/
    "time"
    "crypto/md5"
    "strconv"
    "io"

)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") // write data to response
}
type UserData struct{
    Username string
    Password string
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("login.html")
        t.Execute(w, token)
    } else {
        r.ParseForm()
        token := r.Form.Get("token")
        if token != "" {
            // check token validity
        } else {
            // give error if no token
        }
     //hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 8)
        fmt.Println("username length:", len(r.Form["username"][0]))
        fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // print in server side
        fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
        template.HTMLEscape(w, []byte(r.Form.Get("username"))) // respond to client

        
    }
}
type MyMux struct{

}
func (*MyMux) ServeHTTP(w http.ResponseWriter,r *http.Request){
  if r.URL.Path == "/"{
    sayhelloName(w,r)
    return

  }
  if r.URL.Path == "/login"{
    login(w,r)
    return 
  }

}

func main() {
    mux := &MyMux{}
    err := http.ListenAndServe(":8080", mux) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
