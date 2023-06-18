package url

import (
	"net/url"
	"time"
)

// tipo url para armazenar o identificador curto
type Url struct {
	Id      string    `json:"id"`
	Created time.Time `json:"created"`
	Dest    string    `json:"dest"`
}

func SearchOrCreateAnewUrl(url string) (url.URL, bool, error) {

}
