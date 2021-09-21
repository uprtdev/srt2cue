package main

import (
	"log"
	"time"
)

func ReduceSegments(data []Segment, timeGapSec int) []Segment {
	result := make([]Segment, 0)

	if (len(data)) > 0 {
		firstSegment := &data[0]
		lastSegment := &data[0]
		timeGap := time.Duration(timeGapSec) * time.Second

		for i := range data {
			segment := &data[i]
			if segment.GroupName != lastSegment.GroupName || segment.Moment.Sub(lastSegment.Moment) > timeGap {
				finalSegment := *firstSegment
				finalSegment.MomentText = firstSegment.MomentText + " - " + lastSegment.MomentText
				result = append(result, finalSegment)
				firstSegment = segment
			}
			lastSegment = segment
		}
		log.Printf("Reduced data to %v tracks", len(result))
	} else {
		log.Printf("Got no data, nothing to do here")
	}
	return result
}
