package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func hostname() string {
	hostname, err := os.Hostname()
	if err == nil {
		return hostname
	}
	return "localhost"
}

func ip_addr() string {
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get("http://169.254.169.254/latest/meta-data/public-ipv4")
	if err != nil {
		fmt.Println("Can't get public IP via aws meta data")
		return "127.0.0.1"
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println("Can't get responce Body")
			return "127.0.0.1"
		}
		bodyString := string(bodyBytes)
		return bodyString
	}
	return "Status is Not OK"
}

func main() {
	fmt.Println("Server starting...")
	hostname := hostname()
	fmt.Println("Hostname: " + hostname)
	ip_addr := ip_addr()
	fmt.Println("IP Addr: " + ip_addr)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, ("<p>URL Path: " + r.URL.Path + "</p>"))
		fmt.Fprintf(w, ("<h1>Host: " + hostname + "</h1>"))
		fmt.Fprintf(w, ("<h1>IP Addr: " + ip_addr + "</h1>"))
		fmt.Println(hostname)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
