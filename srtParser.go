package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

// very simple state machine enum for parsing
type State int64

// we will need it for a tiny hack to return the number of milliseconds from the beginning
var refTime, _ = time.Parse("03:04", "00:00")

const SrtTimeDelimiter = " --> "

const (
	Start State = iota
	TimeCode
	Date
	Group
	Other
)

func parseDate(text string) time.Time {
	moment, err := time.Parse("2006/01/2 15:04:05", text)
	if err != nil {
		log.Fatal(err)
	}
	return moment
}

func parseTimeCodeText(text string) time.Duration {
	moment, err := time.Parse("15:04:05,000", text)
	if err != nil {
		log.Fatal(err)
	}
	return moment.Sub(refTime)
}

func parseTimeCodeLine(text string) (time.Duration, time.Duration) {
	times := strings.Split(text, SrtTimeDelimiter)
	if len(times) != 2 {
		log.Fatal("Incorrect file format, can't parse this: %s", text)
	}
	start := parseTimeCodeText(times[0])
	end := parseTimeCodeText(times[1])
	return start, end
}

func ParseSrtFile(filename string) []Segment {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Printf("Opened %s", filename)

	result := make([]Segment, 0)
	var currentSegment Segment
	scanner := bufio.NewScanner(file)

	state := Start
	for scanner.Scan() {
		switch state {
		case Start:
			currentSegment = Segment{}
			state = TimeCode
		case TimeCode:
			currentSegment.Start, currentSegment.End = parseTimeCodeLine(scanner.Text())
			state = Date
		case Date:
			currentSegment.MomentText = scanner.Text()
			currentSegment.Moment = parseDate(scanner.Text())
			state = Group
		case Group:
			currentSegment.GroupName = scanner.Text()
			state = Other
		case Other:
			// end of block?
			if scanner.Text() == "" {
				result = append(result, currentSegment)
				state = Start
			}
			// ignore everything else
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if state == Other { // if the file ends not with an empty line
		result = append(result, currentSegment)
	} else if state != Start { // if we are not in Other or Start state
		log.Fatal("Unexpected end of file")
	}

	log.Printf("Loaded and parsed %v segments", len(result))
	return result
}
