package hint

import "fmt"
import (
	"bytes"
	"encoding/xml"
)

// Reporter defines interface that should be implemented to generate a report.
type Reporter interface {
	Collect(problems []Problem)
	Flush() (string, error)
}

// PlainReporter defines reporter that output problems as plain text
type PlainReporter struct {
	buf bytes.Buffer
}

// Collect receives problems for further report generation
func (r *PlainReporter) Collect(ps []Problem) {
	for _, p := range ps {
		r.buf.WriteString(fmt.Sprintf("%s:%v: %s\n", p.File, p.Position, p.Text))
	}
}

// Flush outputs collected problems one by one in plain text
// Expect no errors here
func (r *PlainReporter) Flush() (report string, err error) {
	report = r.buf.String()
	r.buf.Reset()

	return
}

const (
	checkstyleSeverityIgnore  = "ignore"
	checkstyleSeverityInfo    = "info"
	checkstyleSeverityWarning = "warning"
	checkstyleSeverityError   = "error"
)

// checkstyleReporter produces reports in XML Checkstyle format
type checkstyleReporter struct {
	indent   bool // whether to produce pretty report with indent
	problems map[string][]Problem
}

type checkstyleReport struct {
	XMLName xml.Name         `xml:"checkstyle"`
	Version float64          `xml:"version,attr"`
	Files   []checkstyleFile `xml:"file"`
}

type checkstyleFile struct {
	Filename string              `xml:"name,attr"`
	Messages []checkstyleMessage `xml:"error"`
}

// checkstyleReporter produces reports in XML Checkstyle format
type checkstyleMessage struct {
	Line     int    `xml:"line,attr"`
	Column   int    `xml:"column,attr"`
	Severity string `xml:"severity,attr"`
	Message  string `xml:"message,attr"`
}

// NewCheckstyleReporter creates initialized Checkstyle report generator
func NewCheckstyleReporter(indent bool) *checkstyleReporter {
	return &checkstyleReporter{problems: map[string][]Problem{}, indent: indent}
}

// Collect receives problems for further report generation
func (r *checkstyleReporter) Collect(ps []Problem) {
	for _, p := range ps {
		if _, ok := r.problems[p.Position.Filename]; !ok {
			// In most cases we will call this method once for the same file. Let's allocate memory for it
			r.problems[p.Position.Filename] = make([]Problem, 0, len(ps))

			fmt.Println(p.Position.Filename)
		}
		r.problems[p.File] = append(r.problems[p.File], p)
	}
}

// Flush produces XML in checkstyle format. All the problems are grouped by files
// Returned error indicates that something wrong happened to XML marshalling process
func (r *checkstyleReporter) Flush() (result string, err error) {
	report := checkstyleReport{
		Version: 4.3,
		Files:   make([]checkstyleFile, 0, len(r.problems)),
	}

	for filename, ps := range r.problems {
		file := checkstyleFile{
			Filename: filename,
			Messages: make([]checkstyleMessage, len(ps)),
		}

		for i, p := range ps {
			message := checkstyleMessage{
				Line:     p.Position.Line,
				Column:   p.Position.Column,
				Severity: r.confidenceToSeverity(p.Confidence),
				Message:  p.Text,
			}

			file.Messages[i] = message
		}

		report.Files = append(report.Files, file)

	}

	var resBytes []byte
	if r.indent {
		resBytes, err = xml.MarshalIndent(report, "", "    ")
	} else {
		resBytes, err = xml.Marshal(report)
	}

	if resBytes != nil {
		result = string(resBytes)
	}

	r.problems = map[string][]Problem{}

	return
}

func (r *checkstyleReporter) confidenceToSeverity(c float64) string {
	if c < .25 {
		return checkstyleSeverityIgnore
	} else if c < .5 {
		return checkstyleSeverityInfo
	} else if c < .75 {
		return checkstyleSeverityWarning
	}
	return checkstyleSeverityError
}
