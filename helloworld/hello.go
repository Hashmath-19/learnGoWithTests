package helloworld

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "Worldw"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Hash"))
}
