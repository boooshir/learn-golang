package main

import (
	"flag"
	"fmt"
	"go-channels/deep"
	"go-channels/gemini"
	"go-channels/grok"
)

func main() {
	fromPtr := flag.String("from", "grok", "the learning source")

	flag.Parse()

	fromSource := *fromPtr

	fmt.Println("from", fromSource)

	switch fromSource {
	case "gemini":
		gemini.FromGemini()
	case "deep":
		deep.DeepRun()
	case "deep-jobs":
		deep.RunJobs()
	default:
		grok.FromGrok()
	}
}
