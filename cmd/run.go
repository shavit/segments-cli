package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	"github.com/shavit/segments-cli/stream"
)

func checkError(err error) {
	if err != nil {
		print("\u001b[31m")
		println("Error:", err.Error())
		print("\u001b[0m")
		os.Exit(1)
	}
}

func main() {
	var fPath = flag.String("p", "lib/silence_1.xml", "XML path")
	var fMs = flag.Int("ms", 1000, "Silence duration")
	var fMax = flag.Int("max", 2000, "Maximum segment duration")
	var fSplit = flag.Int("split", 800, "Split duration")
	var fS = flag.Bool("s", false, "Silence the program")
	var fO = flag.String("o", "", "Output path")
	var fH = flag.Bool("h", false, "Help menu")
	flag.Parse()

	// Input validation
	if *fH == true {
		printHelp()
		os.Exit(0)
	}

	if len(os.Args) <= 4 {
		printHelp()
		os.Exit(1)
	}

	if *fS == true && *fO == "" {
		print("\u001b[31m")
		println(`
No output option was given

Please remove the silence option or choose an output path
`)
		print("\u001b[0m")
		os.Exit(2)
	}

	xmlFile, err := os.Open(*fPath)
	checkError(err)

	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)
	silences, err := stream.DecodeSilences(b)
	checkError(err)

	segments := stream.CreateChapters(silences, *fMs, *fMax, *fSplit)
	out, err := json.Marshal(segments)
	checkError(err)

	// Output options

	if *fS == false {
		os.Stdout.Write(out)
	}
	if *fO != "" {
		err = ioutil.WriteFile(*fO, out, 0644)
		checkError(err)
		print("\u001b[32m")
		println("\nFile was successfully written to", *fO)
		print("\u001b[0m")
	} else {
		print("\u001b[32m")
		println("\n\nDone")
		print("\u001b[0m")
	}
}

func printHelp() {
	print("\u001b[33m")
	println(`
    Usage: run path ms max split [OPTIONS]

    Arguments:
      p - Path to the XML file
      ms - Silence duration in milliseconds, that indicates chapter transition
      max - Maximum duration of a segment in milliseconds
      split - Silence duration in milliseconds, to be used for split

    Options:
      o - Output path for the JSON file
      s - Silence the output
`)
	print("\u001b[0m")
}
