package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	fmt.Fprintf(w, "Good morning Beesh, today is: %s ", dt.Format("01-02-2006 Mon"))
}
func main() {
	// fmt.Println("hi beesh")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
