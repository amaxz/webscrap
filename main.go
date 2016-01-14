package main

import (
	"log"
	"os"
	"webscrap/scrap"
	"time"
	"math/rand"
	"runtime"
	"flag"
	"fmt"
	"bufio"
)

var OutputFileName = flag.String("o", "webscrap.log", "Output file name")
var InputFileName = flag.String("f", "", "Input file name (alternative with environment $WOEGO_WEBSCRAP_FILE)")
var SleepSeconds = flag.Int("s", 10, "Minimum sleep duration second")
var keywords []string = make([]string, 10)

func init() {

	flag.Parse()

	logFile, err := os.OpenFile(*OutputFileName, os.O_RDWR | os.O_CREATE | os.O_SYNC | os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	//defer logFile.Close()

	//log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)

}

func ScrapKeyword(keyword string) []scrap.Fetch {
	return []scrap.Fetch{scrap.LoadSuning(keyword), scrap.LoadJD(keyword)}
	//time.Sleep(time.Duration(1) * time.Second)
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	path := *InputFileName
	if path == "" {
		if path = os.Getenv("WOEGO_WEBSCRAP_FILE"); path != "" {
			if _, err := os.Stat(path); err != nil {
				fmt.Fprintf(os.Stderr, "Usage: %s [-o <path>] [-f <path> | <keyword>...]\n", os.Args[0])
				flag.PrintDefaults()
				return
			}
		}
	}

	if path != "" {
		file, _ := os.OpenFile(path, os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm)
		defer file.Close()

		re := bufio.NewReader(file)
		size := 10
		for i := 0;; i ++ {
			linebyte, _, err := re.ReadLine()
			if err != nil {
				break
			}
			if string(linebyte) != "" {
				keywords[i % size] = string(linebyte)
				if (i + 1) % size == 0 {
					ScrapList(keywords)
					keywords = make([]string, size)
				}
			}
		}
		ScrapList(keywords)
	}

	if flag.NArg() > 0 {
		keywords = os.Args[1:]
		ScrapList(keywords)
	}
}

func ScrapList(keywords []string) []int {

	status := make([]int, len(keywords))
	for index, key := range keywords {

		key = scrap.FormatKey(key)
		if key == "" {continue}

		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		tmin := *SleepSeconds + random.Intn(*SleepSeconds)

		fetchs := ScrapKeyword(key)
		for _, v := range fetchs {
			if v.Status == 404 {
				status[index] += 1
			}
		}

		if index != len(keywords) - 1 {
			time.Sleep(time.Duration(tmin) * time.Second)
		}
	}
	return status
}
