package transformations

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/transforms/stats"
)

func ExtractFn(line string, emit func(string)) {
	wordRE := regexp.MustCompile(`[a-zA-Z]+('[a-zA-Z]+)?`)
	words := wordRE.FindAllString(line, -1)
	for _, word := range words {
		emit(strings.ToLower(word))
	}
}

func CountWords(s beam.Scope, words beam.PCollection) beam.PCollection {
	return stats.Count(s, words)
}

func FilterDuplicates(s beam.Scope, counts beam.PCollection) beam.PCollection {
	return beam.ParDo(s, func(word string, count int, emit func(string)) {
		if count > 1 {
			emit(fmt.Sprintf("%s: %d", word, count))
		}
	}, counts)
}