package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

const (
	VERSION_MAJOR = "0"
	VERSION_MINOR = "1"
	VERSION_PATCH = "0"

	ISO8601 = "2006-01-02T15:04:05-0700" // ISO-8601 date-time formatting string.
)

// Versioning detail information, set during build phase.
var (
	BuildTime       time.Time
	BuildTimeStr    string
	BuildCommitHash string
)

func init() {
	// Parse BuildTimeStr into BuildTime.
	if BuildTimeStr != "" {
		var err error
		if BuildTime, err = time.Parse(ISO8601, BuildTimeStr); err != nil {
			panic(fmt.Sprintf(`Parsing BuildTimeStr="%v" failed: %s`, BuildTimeStr, err))
		}
		BuildTime = BuildTime.UTC() // NB: Otherwise the date is displayed as '+0000 +0000' instead of '+0000 UTC'.
	}
	flag.Usage = func() {
		DisplayVersion(os.Stderr, "\n")
		fmt.Fprintln(os.Stderr, "usage:")
		flag.PrintDefaults()
	}
}

func DisplayVersion(writer io.Writer, delimiter string) {
	fmt.Fprintf(writer, "%v.%v.%v%v", VERSION_MAJOR, VERSION_MINOR, VERSION_PATCH, delimiter)
	fmt.Fprintf(writer, "built on %s%v", BuildTime, delimiter)
	fmt.Fprintf(writer, "git commit %v%v", BuildCommitHash, delimiter)
}
