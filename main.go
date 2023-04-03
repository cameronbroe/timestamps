package main

import (
	"flag"
	"fmt"
	"github.com/cameronbroe/timeutil"
	"time"
)

func main() {
	formatArg := flag.String("format", "%H:%M:%S.%s", "format to display timestamp in: https://github.com/cameronbroe/timeutil#strftime")
	rateArg := flag.Int("rate", 1, "rate in milliseconds in which to display timestamp (ms)")

	flag.Parse()

	ticker := time.NewTicker(time.Millisecond * time.Duration(*rateArg))
	for {
		select {
		case time := <-ticker.C:
			fmt.Printf("%s\n", timeutil.Strftime(&time, *formatArg))
		}
	}
}
