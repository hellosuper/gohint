// Copyright (c) 2013 The Go Authors. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd.

// golint lints the Go source files named on its command line.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/elgris/lint"
)

var reporterName = flag.String("reporter", "plain", "name of reported to generate ouput. Available: plain, checkstyle")
var configFile = flag.String("config", "", "path to file with config")
var config *lint.Config

var reporter lint.Reporter

func main() {
	flag.Parse()

	switch *reporterName {
	case "plain":
		reporter = &lint.PlainReporter{}
	case "checkstyle":
		reporter = lint.NewCheckstyleReporter(true)
	default:
		fmt.Fprintf(os.Stderr, "Unknown reporter '%s'. Available ones: plain, checkstyle\n", *reporterName)
		return
	}

	var err error
	config, err = lint.NewConfig(*configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		return
	}

	for _, filename := range flag.Args() {
		if isDir(filename) {
			lintDir(filename)
		} else {
			lintFile(filename)
		}
	}

	report, err := reporter.Flush()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		return
	}
	fmt.Println(report)
}

func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}

func lintFile(filename string) {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	l := new(lint.Linter)
	ps, err := l.Lint(filename, config, src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v:%v\n", filename, err)
		return
	}
	reporter.Collect(ps)
}

func lintDir(dirname string) {
	filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && strings.HasSuffix(path, ".go") {
			lintFile(path)
		}
		return err
	})
}
