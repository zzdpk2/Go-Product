package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	res := make(map[string]bool)
	resChan := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resChan <- result{
				string: u,
				bool:   wc(u),
			}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		resU := <-resChan
		res[resU.string] = resU.bool
	}
	return res
}
