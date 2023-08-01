package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://ipconfig.io/ip")

	if err != nil {
		fmt.Println(fmt.Errorf("Could not retrieve public IP, check your connection."))
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(fmt.Errorf("Could not read response from ipconfig.io."))
	}

	fmt.Print(string(body))
}
