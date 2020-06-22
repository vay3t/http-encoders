# http-encoders
http encoders utils with go

# Help

```
Usage: ./gohttpencoder -e <encoder> <file-stdin>
  -e string
    	Encoders to use:
    		urle: urlencoder
    		durle: doble urlencoder
    		allurle: all urlencoder
    		alldurle: all doble urlencoder
    		b64: base64
    		hex: hexadecimal
    		unic: unicode encoder
  -h	Show this help
```

# Build 

```bash
git clone https://github.com/vay3t/http-encoders
cd http-encoders
go build http-encoders.go
```
