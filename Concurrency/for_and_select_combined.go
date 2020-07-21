package main

import "fmt"

func getNews(newsChannel chan string) {
	NewsArray := []string{"News Item 1", "Breaking News Incoming"}
	for _, news := range NewsArray {
		newsChannel <- news
	}

	newsChannel <- "Done"
	close(newsChannel)
}

func main() {
	myNewsChannel := make(chan string)

	go getNews(myNewsChannel)

	for {
		select {
		case news := <-myNewsChannel:
			fmt.Println(news)
			if news == "Done" {
				return
			}
		default:
		}
	}

	//another example

	done := make(chan string)

	for _, animal := range []string{"elephant", "lion", "tiger"} {
		select {
		case <-done:
			return
		default:
			fmt.Println(animal)
		}
	}

	fmt.Println("MAIN DONE")
}
