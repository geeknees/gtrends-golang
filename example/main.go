package main

import (
	"fmt"
	"reflect"

	"github.com/geeknees/gtrends-golang"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	r, err := gtrends.TrendsIn("")
	handleError(err, "Failed to get daily searches")
	printItems(r)
}

func handleError(err error, errMsg string) {
	if err != nil {
		log.Fatal(errors.Wrap(err, errMsg))
	}
}

func printItems(items interface{}) {
	ref := reflect.ValueOf(items)

	if ref.Kind() != reflect.Slice {
		log.Fatalf("Failed to print %s. It's not a slice type.", ref.Kind())
	}

	for i := 0; i < ref.Len(); i++ {
		log.Println(ref.Index(i).Interface().(*gtrends.TrendingSearch).Title)
	}

}

func p(a ...interface{}) {
	fmt.Println(a...)
}
