package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const CueIndentation = "	"

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func formatTimeCode(timeCode time.Duration, timeShift int) string {
	minute := int64(timeCode.Seconds()) / 60
	second := int64(timeCode.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", minute, second, timeShift)
}

func SaveCueFile(filename string, data []Segment, timeShift int) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	trackCount := 0
	for i := range data {
		segment := &data[i]
		trackCount++
		_, err = fmt.Fprintf(w, "TRACK %02d AUDIO\n", trackCount)
		check(err)
		_, err = fmt.Fprintf(w, "%sTITLE \"%s\"\n", CueIndentation, segment.MomentText)
		check(err)
		_, err = fmt.Fprintf(w, "%sPERFORMER  \"%s\"\n", CueIndentation, segment.GroupName)
		check(err)
		_, err = fmt.Fprintf(w, "%sINDEX 01 %s\n", CueIndentation, formatTimeCode(segment.Start, timeShift))
		check(err)
	}
	w.Flush()

	log.Printf("Saved %v tracks to %s", trackCount, filename)
}
