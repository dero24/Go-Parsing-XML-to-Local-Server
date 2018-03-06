package main

import (
  "encoding/xml"
  "fmt"
  "io/ioutil"
  "net/http"
)

var s Sitemapindex

type Sitemapindex struct {
  Locations []string `xml:"url>loc"` // Locations is a slice, of a Location type
}

func main() {
  resp, _ := http.Get("https://techcrunch.com/news-sitemap.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  xml.Unmarshal(bytes, &s)
  //fmt.Println(s.Locations)

  http.HandleFunc("/", index_handler)
  http.HandleFunc("/about/", about_handler)
  http.ListenAndServe(":8000", nil) 
}



func index_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Latest techcrunch.com posts!</h1>")
  for _, Location := range s.Locations {
      fmt.Fprintf(w,"<p>\n%s</p>",Location)
   }   
}

func about_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Go web design by Robby DeRosa")
}
