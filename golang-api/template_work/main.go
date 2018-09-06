package main
import (
	"html/template"
	"net/http"
        "log"
        "fmt"
        "github.com/gorilla/mux"
)

type cities struct{
  data   []string
}

func main() {
        r := mux.NewRouter()
        r.HandleFunc("/",welcome)
        r.HandleFunc("/view",get)
        if err := http.ListenAndServe(":80", r); err != nil {
                log.Fatal(err)
         }}

func welcome(w http.ResponseWriter, r *http.Request){
       tmpl,_ := template.ParseFiles("layout.html")
 //  city :=  cities{[]string{"a","s"}}
   m := map[string]interface{}{
    "data":   []string{"w","d", "s"},
}
      tmpl.Execute(w,m)
}

func get(w http.ResponseWriter, r *http.Request){
r.ParseForm()
fmt.Fprintln(w,"the choice is ",r.Form)
}
