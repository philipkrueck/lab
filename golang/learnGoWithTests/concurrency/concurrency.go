package concurrency

type (
	WebsiteChecker func(string) bool
	result         struct {
		url     string
		checked bool
	}
)

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	checkedSites := make(map[string]bool)
	ch := make(chan result)

	for _, url := range urls {
		go func() {
			ch <- result{url, wc(url)}
		}()
	}

	for i := len(urls); i > 0; i-- {
		res := <-ch
		checkedSites[res.url] = res.checked
	}

	return checkedSites
}
