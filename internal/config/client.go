package config

import (
	"net/http"
)

var RequestUrl = "https://www.bing.com/images/create?q=%s&rt=3&FORM=GENCRE"
var RequestResultUrl = "https://www.bing.com/images/create/async/results/"

var Cache = make(map[string]string)

var Client = http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
},
}
