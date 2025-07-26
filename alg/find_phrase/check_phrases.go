package find_phrase

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func FindPhrase(phrase string) {

	var filenames = []string{"test.txt", "text1.txt", "text2.txt"}
	var formatPhrase = strings.ToLower(phrase)

	dataChan := make(chan string)
	ackChan := make(chan bool)

	wg := &sync.WaitGroup{}

	fmt.Println("Start scanning")
	go ScanFiles(filenames, dataChan, ackChan, wg)

	for i := 0; i < len(filenames); i++ {
		wg.Add(1)
		go CheckPhrase(formatPhrase, dataChan, ackChan, wg)
	}

	wg.Wait()

}

func ScanFiles(filenames []string, out chan<- string, ack chan bool, wg *sync.WaitGroup) {
	for _, filename := range filenames {
		Content := ScanFile(filename)
		fmt.Printf("file %s\n", filename)

		out <- Content

		<-ack

	}
	close(out)
}

func CheckPhrase(phrase string, in <-chan string, ack chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	Content := <-in
	result := strings.Count(Content, phrase)

	fmt.Println(result)

	ack <- true

}

func FormatText(text string) string {
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ToLower(text)
	return text
}

func ScanFile(filename string) string {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Read file error: %v\n", err)
	}

	var strFileContent = string(fileContent)
	strFileContent = FormatText(strFileContent)

	return string(strFileContent)
}
