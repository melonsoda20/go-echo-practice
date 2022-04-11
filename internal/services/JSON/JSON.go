package json

import (
	"encoding/json"
	"os"
)

func DecodeToJSON(fileData *os.File) *json.Decoder {
	return json.NewDecoder(fileData)
}
