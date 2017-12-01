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

func getMetaData(meta_data_item string) string {
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get("http://169.254.169.254/latest/meta-data/" + meta_data_item + "/")
	if err != nil {
		fmt.Println("Can't get aws meta data item: " + meta_data_item)
		return "can't get " + meta_data_item
	}
	defer resp.Body.Close() // Close body only if response non-nil
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println("Can't get responce Body for meta data item: " + meta_data_item)
			return "can't get " + meta_data_item
		}
		bodyString := string(bodyBytes)
		return bodyString
	}
	return "Status is Not OK for " + meta_data_item
}

func gePublicIP() string {
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get("http://169.254.169.254/latest/meta-data/public-ipv4")
	if err != nil {
		fmt.Println("Can't get public IP via aws meta data")
		return "127.0.0.1"
	}
	defer resp.Body.Close() // Close body only if response non-nil
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
	meta_data_items := []string{
		"ami-id",
		"hostname",
		"instance-id",
		"instance-type",
		"local-hostname",
		"local-ipv4",
		"public-ipv4",
	}

	meta_data_items_map := make(map[string]string)

	for _, item := range meta_data_items {
		meta_data_items_map[item] = getMetaData(item)
	}
	fmt.Println("Server starting...")
	hostname := hostname()
	fmt.Println("Hostname: " + hostname)
	gePublicIP := gePublicIP()
	fmt.Println("IP Addr: " + gePublicIP)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, ("<p>URL Path: " + r.URL.Path + "</p>"))
		fmt.Fprintf(w, ("<h1>Host: " + hostname + "</h1>"))
		fmt.Fprintf(w, ("<h1>IP Addr: " + gePublicIP + "</h1>"))
		for _, item := range meta_data_items {
			fmt.Fprintf(w, ("<h1>" + item + ": " + meta_data_items_map[item] + "</h1>"))
		}
		fmt.Println(hostname)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
