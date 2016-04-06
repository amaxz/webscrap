## Web Scrap 

Webscrap is a price fetcher of commodity sells on e-commerce website like `JD.com`, `Suning.com`, `Tmall.com`, etc

### How to build

* win32

        env GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -ldflags -s -o webscrap-win32.exe webscrap

* linux

        env GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -ldflags -s -o webscrap-linux webscrap

* OS X

        env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags -w -o webscrap-darwin webscrap

<!--
env GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -ldflags -s -o webscrap-win32.exe webscrap
env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags -w -o webscrap-darwin webscrap
env GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -ldflags -s -o webscrap-linux webscrap
-->

### How to run

Usage: `./webscrap [-o <path>] [-f <path> | <keyword>...]`
    
* `-f path`

    Input file name (alternative with environment `$WOEGO_WEBSCRAP_FILE`), the input file is a plain text file of keywords which splited with line seperator '\n'

* `-o path`

    Output file name (default "./output/S`${UnixTime}`.txt"), search result output with json format line text prefixed with timestamp  

* `-s number`

    Minimum sleep duration second (default 10)
    
* `<keyword>...`

    Query keyword, ex. `./webscrap "iPhone 5s" "SONY Z3"`
