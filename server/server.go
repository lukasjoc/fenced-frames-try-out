package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port string = ":4000"
)

func crabHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Supports-Loading-Mode", "fenced-frame")
		fmt.Fprintln(w, `
<h1> we got jebaited </h1>
<div>
<!-- <iframe
    width=840
    height=420
    src="https://youtube.com/embed/LDU_Txk06tM"
    style="border: 0"
    fetchpriority=low
    allowfullscreen>
</iframe> -->
</div>
<script> <!--  async=""> -->
const bodyElement = document.querySelector("body")
iframeElement = document.createElement("iframe")
iframeElement.setAttribute("title", "crab rave")
iframeElement.setAttribute("width", 696)
iframeElement.setAttribute("height", 420)
iframeElement.setAttribute("src", "https://youtube.com/embed/LDU_Txk06tM")
iframeElement.setAttribute("fetchpriority", "high")
iframeElement.allowfullscreen = true
console.log("Setting Iframe: " + iframeElement.title + " into body")
bodyElement.append(iframeElement)
</script>
`)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/crab", crabHandler())
	log.Fatal(http.ListenAndServe(port, mux))
}
