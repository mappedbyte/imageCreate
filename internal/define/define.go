package define

import (
	"encoding/json"
	"errors"
	"image-Designer/internal/task"
	"io/ioutil"
	"os"
)

var ConfigName = "imageDesigner.cnf"

func init() {
	dir, _ := os.Getwd()
	_, err := ioutil.ReadFile(dir + string(os.PathSeparator) + ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		conf := new(Config)
		conf.Cookie = ""
		conf.ProxyEnable = false
		conf.ProxyUrl = ""
		marshal, _ := json.Marshal(conf)
		ioutil.WriteFile(dir+string(os.PathSeparator)+ConfigName, marshal, 0666)
	}
	task.ClearIdCache()
}

type Config struct {
	ProxyEnable bool   `json:"proxyEnable"`
	ProxyUrl    string `json:"proxyUrl"`
	Cookie      string `json:"cookie"`
}
