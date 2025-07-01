// package main

import (
	"fmt"
	"net/url"
)

func main() {

	//we can convert to url object and use method to it
	//we also can create custom url (i skipped this )
	raw_url := "http://example.com:8080/path/to/resource#amrfraf"

	url, error := url.Parse(raw_url)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("fragement: ", url.Fragment)
	fmt.Println("host: ", url.Host)
	fmt.Println("omithost: ", url.OmitHost)
	fmt.Println("path: ", url.Path)
	fmt.Println("scheme: ", url.Scheme)

}
