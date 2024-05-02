package main

import "net/http"
import "fmt"
func main() {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get("https://distopia.savi2w.workers.dev/")

	if err != nil{
	fmt.Println("Error",err)
	}

	
	fmt.Println(resp.Header.Get("distopia"))
	
	
}
