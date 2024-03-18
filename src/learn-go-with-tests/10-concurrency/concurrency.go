package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency#write-enough-code-to-make-it-pass
