package function

import (
	"encoding/json"

	"github.com/mlabouardy/9gag"
)

// Handle a serverless request
func Handle(req []byte) string {
	gag9 := gag9.New()
	memes := gag9.FindByTag(string(req))
	rawJson, _ := json.Marshal(memes)
	return string(rawJson)
}
