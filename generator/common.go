package generator

import (
	"encoding/json"

	"github.com/douchunrong/golang-struct-to-elastic-mapping/mapping"
)

func wrap(properties []byte) ([]byte, error) {
	return json.MarshalIndent(&mapping.Index{
		Mappings: mapping.IndexMappings{
			Properties: properties,
		},
	}, "", "  ")
}

func lastIndex(length, index int) bool {
	return length-1 == index
}
