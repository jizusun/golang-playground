package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	url := os.Args[2]
	wg.Add(limit)

	for i := 0; i < limit; i++ {
		go get(i, &wg, url)
	}
	wg.Wait()
	fmt.Println("end")
}

func get(num int, wg *sync.WaitGroup, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(url, resp.StatusCode, len(body))
	wg.Done()
}
