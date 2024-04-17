package transformations

import (
	"testing"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/passert"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/ptest"
)

func TestExtractFn(t *testing.T) {
	beam.Init()

	p, s := beam.NewPipelineWithRoot()
	lines := beam.Create(s, "Hello world, hello Beam.")
	words := beam.ParDo(s, ExtractFn, lines)

	expected := []string{"hello", "world", "hello", "beam"}
	assert := beam.CreateList(s, expected)
	passert.Equals(s, words, assert)

	if err := ptest.Run(p); err != nil {
		t.Errorf("Failed with error: %v", err)
	}
}

func TestFilterDuplicates(t *testing.T) {
	beam.Init()

	p, s := beam.NewPipelineWithRoot()
	words := beam.Create(s, "hello", "world", "hello", "beam")
	counts := CountWords(s, words)
	duplicates := FilterDuplicates(s, counts)

	expected := []string{"hello: 2"}
	assert := beam.CreateList(s, expected)
	passert.Equals(s, duplicates, assert)

	if err := ptest.Run(p); err != nil {
		t.Errorf("Failed with error: %v", err)
	}
}