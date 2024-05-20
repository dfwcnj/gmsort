package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomstrings() {
	for _ = range 1 << 20 {
		fmt.Println(randSeq(32))
	}
}

func randomuints() {
	for _ = range 1 << 20 {
		fmt.Println(rand.Uint64())
	}
}

func randomdates(format string) {
	var mod = int64(306327747)
	for _ = range 1 << 20 {
		ri := rand.Int63() % mod
		tm := time.Unix(int64(ri), int64(0))

		switch format {
		case "DateTime":
			fmt.Println(tm.Format(time.DateTime))
		case "Layout":
			fmt.Println(tm.Format(time.Layout))
		case "RubyDate":
			fmt.Println(tm.Format(time.RubyDate))
		case "UnixDate":
			fmt.Println(tm.Format(time.UnixDate))
		case "RFC3339":
			fmt.Println(tm.Format(time.RFC3339))
		}
	}
}

func main() {
	var dtype string
	var format string
	flag.StringVar(&dtype, "datatype", "string", "type of data to sort")
	flag.StringVar(&format, "format", "RFC3339", "if dtype is datetime, what date format")
	flag.Parse()

	if dtype == "string" {
		randomstrings()
	} else if dtype == "int" {
		randomuints()
	} else if dtype == "datetime" {
		randomdates(format)
	} else {
		log.Fatal("only string, int, or datetime")
	}
}
