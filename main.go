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
	"strconv"
	"strings"
	"path"
)

var OutputFileName = flag.String("o", "./output/S" + strings.ToUpper(strconv.FormatInt(time.Now().Unix(), 10)) + ".txt", "Output file name")
var InputFileName = flag.String("f", "", "Input file name (alternative with environment $WOEGO_WEBSCRAP_FILE)")
var SleepSeconds = flag.Int("s", 10, "Minimum sleep duration second")

var keywords []string = make([]string, 10)

func init() {

	flag.Parse()

	if _, err := os.Stat(*OutputFileName); err != nil {
		os.Mkdir(path.Dir(*OutputFileName), 0774)
	}
	logFile, err := os.OpenFile(*OutputFileName, os.O_RDWR | os.O_CREATE | os.O_SYNC | os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	//defer logFile.Close()

	//log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetOutput(bufio.NewWriterSize(logFile, 100))
	log.SetFlags(log.Ldate | log.Ltime)

}

func Load(fetcher scrap.Fetcher, task *scrap.Task) {
	fetcher.Load(task)
}

func ScrapKeyword(keyword string) []scrap.Task {
	jd := scrap.Task{ Keyword: keyword, Src: scrap.JD, Fetcher: scrap.JdFetcher{}}
	sn := scrap.Task{ Keyword: keyword, Src: scrap.SUNING, Fetcher: scrap.SuningFetcher{}}
	tmall := scrap.Task{ Keyword: keyword, Src: scrap.TMALL, Fetcher: scrap.TmallFetcher{}}

	return []scrap.Task{tmall, jd, sn}
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	filename := *InputFileName
	if filename == "" {
		if filename = os.Getenv("WOEGO_WEBSCRAP_FILE"); filename != "" {
			if _, err := os.Stat(filename); err != nil {
				fmt.Fprintf(os.Stderr, "Usage: %s [-o <path>] [-f <path> | <keyword>...]\n", os.Args[0])
				flag.PrintDefaults()
				return
			}
		}
	}

	if filename != "" {
		file, _ := os.OpenFile(filename, os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm)
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

		task := ScrapKeyword(key)
		if index > 0 {
			time.Sleep(time.Duration(tmin) * time.Second)
		}
		for _, t := range task {
			fmt.Printf(scrap.KEYLOG_FORMAT, t.Src, t.Keyword)
			Load(t.Fetcher, &t)
			if t.Status == 404 {
				status[index] += 1
			} else {
				for count, item := range t.Items {
					fmt.Printf(scrap.ITEMLOG_FORMAT, count + 1, item.Price, item.Vendor, item.Title, item.Url)
					log.Println(scrap.JsonString(item))
				}
			}
		}
	}
	return status
}
