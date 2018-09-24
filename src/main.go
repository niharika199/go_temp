package main
import (
	"html/template"
	"net/http"
        "log"
        "github.com/gorilla/mux"
)

type city struct{
  Data   []string //the name of the field should start with Capital in order to access in other templates
  Tmp string
}

type temp struct{
Tmp string
}

func main() {
        r := mux.NewRouter()
        r.HandleFunc("/",welcome)
        r.HandleFunc("/hey",hello)
        if err := http.ListenAndServe(":80", r); err != nil {
                log.Fatal(err)
         }}

func welcome(w http.ResponseWriter, r *http.Request){
        t := template.Must(template.ParseGlob("tmpls/*"))
       // C:= city{Data:[]string{"Hyderabad","Chennai","Banglore","Mumbai","Trivandrum","Delhi","Bhubaneshwar"}}
        t.ExecuteTemplate(w,"layout.tmpl",city{Tmp:"body",Data:[]string{"Hyderabad","Chennai","Banglore","Mumbai","Trivandrum","Delhi","Bhubaneshwar"}})
}

func hello(w http.ResponseWriter, r *http.Request){
        t := template.Must(template.ParseGlob("tmpls/*"))
   //   D:= city{Data:[]string{"India","London","US","Canada"}}
        t.ExecuteTemplate(w,"layout.tmpl",city{Tmp:"states",Data:[]string{"India","London","US","Canada"}})
}


