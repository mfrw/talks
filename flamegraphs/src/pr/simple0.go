package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime/pprof"
	"strings"
	"time"
)

var (
	pr           = flag.String("profile", "", "Filename for profile")
	infile       = flag.String("in", "big.txt", "Input file")
	outfile      = flag.String("out", "", "Output file")
	regex        = flag.String("regex", "H([A-Za-z]+)o", "Regualr exp to search")
	totalMatches = 0
)

func trackTime(s time.Time, msg string) {
	e := time.Since(s)
	fmt.Println(msg, ":", e)
}

func main() {
	flag.Parse()
	defer trackTime(time.Now(), "MAIN")

	if *pr != "" {
		f, err := os.Create(*pr)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	/* IO logistics */
	var in, out *os.File
	var err error
	in, err = os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	if *outfile != "" {
		out, err = os.Create(*outfile)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
	} else {
		out = os.Stdout
	}
	/* IO logistics end */

	err = process_file(in, out)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total Found:", totalMatches)
}

func process_file(in, out *os.File) error {
	file, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	temp := strings.Split(string(file), "\n")
	for _, line := range temp {
		dump_match(line, out)
	}
	return nil
}

func dump_match(line string, out *os.File) error {
	re := regexp.MustCompile(*regex)
	match := re.FindStringSubmatch(line)
	if len(match) >= 1 {
		for _, v := range match {
			out.WriteString(v + "\n")
			totalMatches++
		}
	}
	return nil
}
