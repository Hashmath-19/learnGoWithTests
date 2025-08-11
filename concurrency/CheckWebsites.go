package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// Because the only way to start a goroutine is to put go in front of a function call,
		// we often use anonymous functions when we want to start a goroutine
		// Anonymous functions have a number of features which make them useful, two of which
		// we're using above. Firstly, they can be executed at the same time that they're declared
		// - this is what the () at the end of the anonymous function is doing.
		// Secondly they maintain access to the lexical scope in which they are defined -
		// all the variables that are available at the point when you declare the anonymous
		// function are also available in the body of the function.

		// Welcome to concurrency: when it's not handled correctly it's hard to predict
		// what's going to happen. Don't worry - that's why we're writing tests, to help us know
		// when we're handling concurrency predictably.
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
