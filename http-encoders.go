package main

import (
	"bufio"
	"bytes"
	b64 "encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
)

func main() {

	encoder := flag.String("e", "", encodersList())
	welp := flag.Bool("h", false, "Show this help")

	flag.Parse()

	if *encoder != "" {
		switch *encoder {
		case "urle":
			conv2urlenc()
		case "durle":
			conv2dobleurlenc()
		case "allurle":
			conv2allurlenc()
		case "alldurle":
			conv2alldobleurlenc()
		case "b64":
			conv2base64()
		case "hex":
			conv2hex()
		case "unic":
			conv2unicode()
		default:
			usage()
		}

	} else if *welp == true {
		usage()
	} else {
		usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s -e <encoder> <file-stdin>", os.Args[0])
	fmt.Println()
	flag.PrintDefaults()
}

func encodersList() string {
	c := `Encoders to use:
	urle: urlencoder
	durle: doble urlencoder
	allurle: all urlencoder
	alldurle: all doble urlencoder
	b64: base64
	hex: hexadecimal
	unic: unicode encoder`
	return c
}

func openFile() *bufio.Scanner {
	if len(os.Args) == 4 {
		readFile, err := os.Open(os.Args[3])
		print(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(readFile)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		return scanner
	} else {
		readFile := os.Stdin
		scanner := bufio.NewScanner(readFile)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		return scanner
	}
}

func urlencoder(texto string) string {
	return url.QueryEscape(texto)
}

func dobleurlencoder(texto string) string {
	return url.QueryEscape(url.QueryEscape(texto))
}

func allurlencoder(texto string) string {
	var b bytes.Buffer
	for _, c := range texto {
		c := hex.EncodeToString([]byte(string(c)))
		b.WriteString("%")
		b.WriteString(c)
	}
	return b.String()
}

func alldobleurlencoder(texto string) string {
	return allurlencoder(allurlencoder(texto))
}

func base64(texto string) string {
	return b64.StdEncoding.EncodeToString([]byte(texto))
}

func hexa(texto string) string {
	var b bytes.Buffer
	for _, c := range texto {
		c := hex.EncodeToString([]byte(string(c)))
		b.WriteString("\\x")
		b.WriteString(c)
	}
	return b.String()
}

func unicode(texto string) string {
	var b bytes.Buffer
	for _, c := range texto {
		if c < 128 {
			b.WriteString(string(c))
		} else {
			b.WriteString("\\u")
			b.WriteString(strconv.FormatInt(int64(c), 16))
		}
	}
	return b.String()

}

func conv2urlenc() {
	data := openFile()
	for data.Scan() {
		encoded := urlencoder(data.Text())
		fmt.Println(encoded)
	}
	if err := data.Err(); err != nil {
		log.Println(err)
	}
}

func conv2dobleurlenc() {
	data := openFile()
	for data.Scan() {
		encoded := urlencoder(data.Text())
		fmt.Println(encoded)
	}
	if err := data.Err(); err != nil {
		log.Println(err)
	}
}

func conv2allurlenc() {
	data := openFile()
	for data.Scan() {
		encoded := allurlencoder(data.Text())
		fmt.Println(encoded)
	}
	if err := data.Err(); err != nil {
		log.Println(err)
	}
}

func conv2alldobleurlenc() {
	data := openFile()
	for data.Scan() {
		encoded := alldobleurlencoder(data.Text())
		fmt.Println(encoded)
	}
	if err := data.Err(); err != nil {
		log.Println(err)
	}
}

func conv2base64() {
	data := openFile()
	for data.Scan() {
		encoded := base64(data.Text())
		fmt.Println(encoded)
	}
	if err := data.Err(); err != nil {
		log.Println(err)
	}
}

func conv2hex() {
	data := openFile()
	for data.Scan() {
		encoded := hexa(data.Text())
		fmt.Println(encoded)
	}
	if err := data.Err(); err != nil {
		log.Println(err)
	}
}

func conv2unicode() {
	data := openFile()
	for data.Scan() {
		encoded := unicode(data.Text())
		fmt.Println(encoded)
	}
	if err := data.Err(); err != nil {
		log.Println(err)
	}
}
