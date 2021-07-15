package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const (
	POSTS_COUNT = 100
	BASE_URL    = "https://jsonplaceholder.typicode.com/posts/"
)

var (
	wg sync.WaitGroup
)

func getPost(id int) {

	defer wg.Done()

	resp, err := http.Get(BASE_URL + strconv.Itoa(id))

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(body))

}

func main() {

	for postID := 1; postID <= POSTS_COUNT; postID++ {
		wg.Add(1)
		go getPost(postID)
	}

	wg.Wait()
}
