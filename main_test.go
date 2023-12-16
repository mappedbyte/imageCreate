package main

import (
	"fmt"
	"image-Designer/internal/service"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestName(t *testing.T) {

	//一人之下宝儿姐，汉服婚礼红色服饰，淡颜系，清纯感，韩系,背景是中国古代长安街
	//requestMsg := "一人之下宝儿姐，汉服，淡颜系，清纯感，韩系,背景是中国古代长安街"
	//requestMsg := "一人之下宝儿姐，汉服，淡颜系，清纯感，高冷，韩系,背景是中国古代长安街"
	//requestMsg := "一人之下宝儿姐，秘书制服，淡颜系，清纯感，御姐，韩系"

	config, _ := service.ReadConfig()
	Url := "https://www.bing.com/images/create/async/results/1-657aea27004e4180a3c3500267147c8d?q=%E5%A1%9E%E5%B0%94%E8%BE%BE%E5%85%AC%E4%B8%BB%EF%BC%8C%E6%B7%A1%E9%A2%9C%E7%B3%BB%EF%BC%8C%E6%B8%85%E7%BA%AF%E6%84%9F%EF%BC%8C%E9%9F%A9%E7%B3%BB"

	transport := &http.Transport{
		Proxy: func(_ *http.Request) (*url.URL, error) {

			return url.Parse("http://127.0.0.1:7890")
		},
	}
	client := http.Client{

		Transport: transport,
	}

	request, _ := http.NewRequest("GET", Url, nil)
	request.Header.Set("Cookie", config.Cookie)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0")
	do, _ := client.Do(request)
	bytes, _ := ioutil.ReadAll(do.Body)
	fmt.Println(do.Status, do.ContentLength, string(bytes))

}
