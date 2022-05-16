package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	header = `<!DOCTYPE html>
	<html>
		<head>
			<meta http-equiv="content-type" content="text/html;charset=utf-8">
			<title>Markdown Preview Tool</title>
		</head>
		<body>
	`
	footer = `
		</body>
	</html>
	`
)

func main() {
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
