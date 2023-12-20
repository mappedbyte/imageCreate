package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"image-Designer/internal/config"
	"image-Designer/internal/define"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var Config *define.Config

func init() {
	conf, err := ReadConfig()
	if err != nil {
		log.Fatal("配置文件不存在,程序退出:", err)
		return
	}
	if len(conf.Cookie) == 0 {
		log.Fatal("配置文件Cookie不存在,程序退出:", err)
		return
	}

	if conf.ProxyEnable {
		if len(conf.ProxyUrl) == 0 {
			log.Fatal("开启代理,但代理地址为空,程序退出")
			os.Exit(1)
		}
		transport := &http.Transport{
			Proxy: func(_ *http.Request) (*url.URL, error) {
				return url.Parse(conf.ProxyUrl)
			},
		}
		config.Client.Transport = transport
	}
	Config = conf
}

func ReadConfig() (*define.Config, error) {
	dir, _ := os.Getwd()
	file, err := ioutil.ReadFile(dir + string(os.PathSeparator) + define.ConfigName)
	if err != nil {
		return nil, err
	}
	conf := new(define.Config)
	err = json.Unmarshal(file, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func Submit(message string) (string, error) {

	escape := url.QueryEscape(message)
	// 根据你的代码进行处理，提交请求
	requestUrl := fmt.Sprintf(config.RequestUrl, escape)
	request, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("Cookie", Config.Cookie)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Transfer-Encoding", "chunked")
	request.Header.Add("Host", "https://www.bing.com")
	response, err := config.Client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}
	location, err := response.Location()
	if err != nil {
		return "", err
	}
	newRequest, err := http.NewRequest("GET", location.String(), nil)
	newRequest.Header = request.Header
	newRequest.Header.Add("Referer", requestUrl)
	r, err := config.Client.Do(newRequest)
	defer r.Body.Close()
	if err != nil {
		return "", err
	}
	id := location.Query().Get("id")
	q := location.Query().Get("q")
	if len(id) == 0 {
		return "", errors.New("请求失败,请确认网络,或添加代理")
	}
	config.Cache[id] = q
	return id, nil
}

func Result(id string) ([]string, error) {
	q := config.Cache[id]
	if len(q) == 0 {
		return nil, errors.New("id 不存在,请检查")
	}
	requestUrl := config.RequestResultUrl + id + "?q=" + url.QueryEscape(q)
	request, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Cookie", Config.Cookie)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	response, err := config.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if len(string(bytes)) == 0 {
		return nil, errors.New("照片正在生成中,请稍后重试")
	}
	doc, err := html.Parse(strings.NewReader(string(bytes)))
	if err != nil {
		return nil, err
	}
	nodes, err := htmlquery.QueryAll(doc, "//div[@class='img_cont hoff']/img/@src")
	if err != nil {
		return nil, err
	}

	srcArrays := make([]string, 0)
	for _, node := range nodes {
		// 获取节点的值（src 属性的值）
		srcValue := htmlquery.InnerText(node)
		srcValue = strings.ReplaceAll(srcValue, "w=270&h=270&c=6&r=0&o=5&dpr=1.5", "")
		srcArrays = append(srcArrays, srcValue)
	}

	if len(srcArrays) == 0 {
		nodes, err = htmlquery.QueryAll(doc, "//*[@id='gir_async']/a/img/@src")
		for _, node := range nodes {
			// 获取节点的值（src 属性的值）
			srcValue := htmlquery.InnerText(node)
			srcValue = strings.ReplaceAll(srcValue, "w=270&h=270&c=6&r=0&o=5&dpr=1.5", "")
			srcArrays = append(srcArrays, srcValue)
		}
	}

	if len(srcArrays) == 0 {
		return nil, errors.New("包含敏感词汇,已阻止生成")
	}
	//delete(config.Cache, id)
	return srcArrays, nil
}
