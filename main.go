package main

import (
	"log"
	"os"
	"webscrap/scrap"
	"strings"
	"time"
	"math/rand"
	"runtime"
	"flag"
	"io"
)

func init() {
	logFileName := flag.String("log", "webscrap.log", "Log file name")
	flag.Parse()
	logFile, err := os.OpenFile(*logFileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	//defer logFile.Close()

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetFlags(log.Ldate | log.Ltime)
}


func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	args := os.Args[1:]

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for index := 0; index < len(args); index++ {

		keyword := args[index]
		keyword = strings.Replace(keyword, "-", " ", -1)

		scrap.LoadSuning(keyword)

		if index != len(args) - 1 {
			time.Sleep(time.Duration(2) * time.Second)
		}

		scrap.LoadJD(keyword)

		tmin := 12 + r.Intn(12)

		if index != len(args) - 1 {
			time.Sleep(time.Duration(tmin) * time.Second)
		}
	}

}
