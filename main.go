package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func downloadImage(url string) {
	defer wg.Done()
	fmt.Println("Download image started: ", url)
	time.Sleep(1 * time.Second)
	fmt.Println("Download image completed: ",url)
	
}


func main() {
	numWorkers := 10

	resultImages := []string{"http://image1.png","http://image2.png","http://image3.png","http://image4.png","http://image5.png","http://image6.png","http://image7.png","http://image8.png","http://image9.png","http://image10.png","http://image11.png","http://image12.png","http://image13.png","http://image14.png","http://image15.png"}
    fmt.Println(len(resultImages))
	
	wg.Add(len(resultImages))

	stringChan := make(chan string)

	
	for i := 1; i <=numWorkers; i++ {
		go func() {
			
			for str := range stringChan {
				downloadImage(str)
			}
		}()
	}

	
	startTime := time.Now()
	for _, v := range resultImages {
		stringChan <- v
	}

	
	close(stringChan)

	fmt.Println("Time to download image:", time.Since(startTime))
}
