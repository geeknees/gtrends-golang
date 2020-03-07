// Package gtrends is for learning golang
package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "strings"
)

type Entry struct {
  Default *trendingSearchesDays `json:"default" bson:"default"`
}

type trendingSearchesDays struct {
  Searches []*trendingSearchDays `json:"trendingSearchesDays" bson:"trending_search_days"`
}

type trendingSearchDays struct {
  Searches []*TrendingSearch `json:"trendingSearches" bson:"searches"`
}

type TrendingSearch struct {
  Title *SearchTitle `json:"title" bson:"title"`
}

type SearchTitle struct {
  Query string `json:"query" bson:"query"`
}

// TrendsIn retruns google trends json list
func TrendsIn(country string) {
  // TODO
}

func main() {
  resp, err := http.Get("https://trends.google.com/trends/api/dailytrends?geo=JP")
  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  // google api returns not valid json :(
  data := strings.Replace(string(body), ")]}',", "", 1)

  // var entries interface{}
  out := new(Entry)

  if err := json.Unmarshal([]byte(data), out); err != nil {
    log.Fatal(err)
  }

  // for _, data := range entries.(map[string]interface{}) {
  //   var d = data.(map[string]interface{})
  //   fmt.Println(d["trendingSearchesDays"].([]interface{})[0].(map[string]interface{})["trendingSearches"])
  // }

  searches := make([]*TrendingSearch, 0)
  for _, v := range out.Default.Searches {
    for _, k := range v.Searches {
      searches = append(searches, k)
      fmt.Println(k.Title)
    }
  }
}
