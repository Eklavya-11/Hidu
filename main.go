package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

func init() {
	flag.Usage = func() {
		help := []string{
			"Hidu Hidden Params Finder",
			"",
			"[buffers] | hidu",
			"+=======================================================+",
			"",
			" -h              Displays This Help Message",
			"",
			"+=======================================================+",
			"",
		}

		fmt.Fprintf(os.Stderr, strings.Join(help, "\n"))
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup

	// Use a buffered channel to avoid blocking on writing to the channel
	targets := make(chan string, 50)

	// Create a group of workers to process the targets
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for target := range targets {
				if parsedURL, err := url.Parse(target); err == nil {
					params := getParams(parsedURL)
					if params != "ERROR" {
						fmt.Println(params)
					}
				}
			}
		}()
	}

	// Read targets from input and send them to the workers
	for scanner.Scan() {
		targets <- scanner.Text()
	}
	close(targets)

	wg.Wait()
}

func getParams(url string) string {
	
	// Create a new HTTP client with a timeout of 3 seconds
	client := &http.Client{Timeout: 3 * time.Second}
	
	// Send a GET request to the specified URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "ERROR"
	}
	req.Header.Set("Connection", "close")
	resp, err := client.Do(req)
	if err != nil {
		return "ERROR"
	}
	defer resp.Body.Close()
	
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "ERROR"
	}
	
	// Find all input elements in the HTML response that have a name attribute
	r := regexp.MustCompile(`<input[^>]*name=[\"\']?([^\"\'\s>]+)[^>]*>`)
	matches := r.FindAllSubmatch(body, -1)
	
	// Construct a map of parameter names to values
	params := make(map[string]string)
	for _, match := range matches {
		if len(match) >= 2 {
			name := string(match[1])
			params[name] = "hidu"
		}
	}
	
	// Construct the final URL with the new parameters
	u, err := url.Parse(url)
	if err != nil {
		return "ERROR"
	}
	q := u.Query()
	for name, value := range params {
		q.Set(name, value)
	}
	u.RawQuery = q.Encode()
	
	// Return the final URL
	return u.String()
}
