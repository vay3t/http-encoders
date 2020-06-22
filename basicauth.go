package main

import (
	"os"
	b64 "encoding/base64"
	"fmt"
	"bufio"
	"bytes"
	"log"
)

func main(){
	dict1 := os.Args[1]
	dict2 := os.Args[2]

	file1, err := os.Open(dict1)
	if err != nil { log.Fatal(err) }
	defer file1.Close()
	
	scanner1 := bufio.NewScanner(file1)
	scanner1.Split(bufio.ScanLines)

	for scanner1.Scan(){
		texto1 := scanner1.Text()

		file2, err := os.Open(dict2)
		if err != nil { log.Fatal(err) }
		defer file2.Close()

		scanner2 := bufio.NewScanner(file2)
		scanner2.Split(bufio.ScanLines)

		for scanner2.Scan(){
			texto2 := scanner2.Text()
			
			var b bytes.Buffer
			
			b.WriteString(texto1)
			b.WriteString(":")
			b.WriteString(texto2)

			sEnc := b64.StdEncoding.EncodeToString([]byte(b.Bytes()))

			fmt.Println(sEnc)
		}
	}
}
