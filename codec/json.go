package codec

import (
	"encoding/json"
	"log"
)

func JSONString(v any) string {
	bs, err := json.Marshal(v)
	if err != nil {
		log.Println("json marshal error: ", err)
		return ""
	}
	return string(bs)
}
