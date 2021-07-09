package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

const (
	baseURL = "http://ip-api.com/json/"
	fields  = "?fields=status,message,country,city,proxy"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Country string `json:"country"`
	City    string `json:"city"`
	Proxy   bool   `json:"proxy"`
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Where's the IP bro?")
		os.Exit(1)
	}

	for i, a := range os.Args[1:] {
		if net.ParseIP(a) == nil {
			fmt.Printf("Arg %d (%s) is an invalid IP\n", i+1, a)
		} else {
			goIP(a)
		}
	}
	// flags for additional info
}

func goIP(ip string) {
	var responseBody Response

	resp, err := http.Get(baseURL + ip + fields)
	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		os.Exit(1)
	}

	if resp.StatusCode == 200 {
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed reading data from the response %s\n", err)
			os.Exit(1)
		}

		json.Unmarshal(responseData, &responseBody)

		fmt.Printf("City: %s\nCountry: %s\nProxy: %t\n", responseBody.City, responseBody.Country, responseBody.Proxy)
	}
}
