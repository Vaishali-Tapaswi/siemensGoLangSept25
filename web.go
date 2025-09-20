
package main
import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			//fmt.Fprintf(w, "<h1>Index Page</h1>")
			w.Write([]byte("<h1>Hello World !!</h1>"))
			
		})
	fmt.Println("Starting Server on 8080...")
	http.ListenAndServe(":8080", nil)
}
