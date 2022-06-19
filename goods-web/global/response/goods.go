package response

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

type GoodsResponse struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}
