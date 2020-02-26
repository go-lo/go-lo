package golo

import (
	"fmt"
)

// Tagify takes a map of strings to interfaces, and
// turns it into tags which can be used in Responses
func Tagify(m map[string]interface{}) (tags []*ResponseTag) {
	tags = make([]*ResponseTag, len(m))

	i := 0
	for k, v := range m {
		tags[i] = &ResponseTag{
			Key:   k,
			Value: fmt.Sprintf("%v", v),
		}

		i++
	}

	return
}
