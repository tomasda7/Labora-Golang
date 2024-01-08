package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		postRequestToServer(10000,"/api/v1/inc", "Increment")
	}()

	go func() {
		defer wg.Done()
		postRequestToServer(10000,"/api/v1/dec", "Decrement")
	}()

	wg.Wait()

	response, err := http.Get("http://localhost:9090/api/v1/curr")
	if err != nil {
		fmt.Printf("Error trying to GET /api/v1/curr: %v\n", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body",err)
		} else {
			bodyString := string(bodyBytes)
			fmt.Println("/api/v1/curr response body:", bodyString)
		}
	} else {
		fmt.Println("/api/v1/curr response status code:", response.StatusCode)
	}
}


func postRequestToServer(nRequest int, url string, action string) {
	var postWG sync.WaitGroup
	postWG.Add(nRequest)

	path := fmt.Sprintf("http://localhost:9090%s", url)

	for i := 0; i < nRequest; i++ {
		go func() {
			defer postWG.Done()
			res, err := http.Post(path, "", nil)
			if err != nil {
				fmt.Printf("Error trying to POST %s: %v\n", action, err)
				return
			}
			res.Body.Close()
		}()
		time.Sleep(500 * time.Microsecond)
	}
	fmt.Printf("%s in progress...\n", action)

	postWG.Wait()

	fmt.Printf("%s completed\n", action)
}
