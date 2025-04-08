package readwritefile

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadWriteFile() {
	file, err := os.Open("query.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buffer.WriteString("'")
		buffer.WriteString(scanner.Text())
		buffer.WriteString("'")
		buffer.WriteString(",")
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	// fmt.Println(buffer.String())
	if err = ioutil.WriteFile("result.txt", buffer.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
