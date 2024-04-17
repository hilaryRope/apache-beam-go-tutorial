package main

import (
	"apache-beam-go-tutorial/transformations"
	"context"
	"flag"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/direct"
)

var (
	inputFile  = flag.String("input", "input.txt", "File(s) to read.")
	outputFile = flag.String("output", "output.txt", "Output file (required).")
)

func main() {
	flag.Parse()
	beam.Init()

	if *outputFile == "" {
		log.Fatal(context.Background(), "No output provided")
	}

	p := beam.NewPipeline()
	s := p.Root()

	lines := textio.Read(s, *inputFile)
	words := beam.ParDo(s, transformations.ExtractFn, lines)
	counted := transformations.CountWords(s, words)
	duplicates := transformations.FilterDuplicates(s, counted)

	textio.Write(s, *outputFile, duplicates)

	if _, err := direct.Execute(context.Background(), p); err != nil {
		log.Fatalf(context.Background(), "Failed to execute job: %v", err)
	}
}