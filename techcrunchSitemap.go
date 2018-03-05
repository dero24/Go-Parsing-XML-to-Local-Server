package main

import (
  "encoding/xml"
  "fmt"
  "io/ioutil"
  "net/http"
)

type Sitemapindex struct {
  Locations []string `xml:"url>loc"` // Locations is a slice, of a Location type
}

func main() {
  resp, _ := http.Get("https://techcrunch.com/news-sitemap.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  var s Sitemapindex
  xml.Unmarshal(bytes, &s)
  //fmt.Println(s.Locations)
  for _, Location := range s.Locations {
      fmt.Printf("\n%s",Location)
  }
}
