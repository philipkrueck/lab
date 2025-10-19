package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type Cache struct {
	v  map[string]bool
	mu sync.Mutex
}

func (cache *Cache) Set(key string, value bool) {
	cache.mu.Lock()
	cache.v[key] = value
	cache.mu.Unlock()
}

func (cache *Cache) Get(key string) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	return cache.v[key]
}

func Crawl(url string, depth int, fetcher Fetcher, visited *Cache, wg *sync.WaitGroup) {
	if depth == 0 || visited.Get(url) {
		wg.Done()
		return
	}

	visited.Set(url, true)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, visited, wg)
	}
	wg.Done()
}

func main() {
	visited := Cache{v: make(map[string]bool)}

	var wg sync.WaitGroup
	wg.Add(1)

	Crawl("https://golang.org/", 4, fetcher, &visited, &wg)

	wg.Wait()
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (body string, urls []string, err error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", []string{}, fmt.Errorf("not found %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
