package main

import (
	"flag"
	"fmt"
	"log"
)

const Version = "1.0"
const MaxCueFrames = 75

func processFile(inputFile string, outputFile string, timeGap int, timeShift int) {
	originalData := ParseSrtFile(inputFile)
	reducedData := ReduceSegments(originalData, timeGap)
	SaveCueFile(outputFile, reducedData, timeShift)
}

func printBanner() {
	fmt.Println("*** srt2cue v" + Version)
	fmt.Println("*** https://github.com/uprt/srt2cue")
}

func main() {
	printBanner()
	var timeGapArg = flag.Int("t", 5, "Time gap between pieces to split them to separate tracks")
	var timeShiftArg = flag.Int("s", 0, "Time shift for the moment of splitting tracks in CUE frames (0-75)")
	var inputFileArg = flag.String("i", "", "Input file name/path")
	var outputFileArg = flag.String("o", "", "Output file name/path")
	flag.Parse()

	if *timeShiftArg > MaxCueFrames {
		log.Fatal("Incorrect time shift value, can be 0-75")
	}

	if *inputFileArg == "" {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		log.Fatal("Please specify input file!")
	}

	outputFile := *outputFileArg
	if *outputFileArg == "" {
		outputFile = replaceExtension(*inputFileArg, ".cue")
		log.Printf("No output file specified, will save output to %s", outputFile)
	}

	processFile(*inputFileArg, outputFile, *timeGapArg, *timeShiftArg)
}
