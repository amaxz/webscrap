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
)

func init() {
	logFileName := flag.String("log", "scrap.log", "Log file name")
	flag.Parse()
	logFile, err := os.OpenFile(*logFileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0622)
	if err != nil {
		panic(err)
	}
	//defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)
}


func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	args := os.Args[1:]

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for index := 0; index < len(args); index++ {

		keyword := args[index]
		keyword = strings.Replace(keyword, "-", " ", -1)

		log.Printf(scrap.KEYLOG_FORMAT, keyword)
		scrap.LoadJD(keyword)

		tmin := 8 + r.Intn(12)

		if index != len(args) - 1 {
			time.Sleep(time.Duration(tmin) * time.Second)
		}
	}

}
