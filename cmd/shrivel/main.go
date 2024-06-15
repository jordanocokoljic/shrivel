package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jordanocokoljic/shrivel"
)

//go:embed usage.txt
var usageText string

func main() {
	in := flag.String("in", "", "An (optional) path to a file to read text to minify from")
	out := flag.String("out", "", "An (optional) path to write minified text to")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usageText)
		flag.PrintDefaults()
	}

	flag.Parse()

	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		flag.Usage()
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	if *in != "" {
		file, err := os.Open(*in)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		reader = bufio.NewReader(file)
	}

	writer := bufio.NewWriter(os.Stdout)
	if *out != "" {
		file, err := os.Create(*out)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		writer = bufio.NewWriter(file)
	}

	src := make([]rune, 0, 512)

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		src = append(src, r)
	}

	src = src[:shrivel.Sql(src, src)]

	if _, err := writer.WriteString(string(src)); err != nil {
		log.Fatal(err)
	}

	if err := writer.Flush(); err != nil {
		log.Fatal(err)
	}
}
