package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type List struct {
	Torrents map[string]Torrents
}
type Torrents struct {
	Infohash_V1 string
	Up_Limit    int64
	Tracker     string
}

func Get_List(host string) map[string]Torrents {
	var list List
	resp, err := http.Get(host + api_list)
	if err != nil {
		log.Println(err.Error())
		return list.Torrents
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal([]byte(body), &list)
	return list.Torrents
}

func Set_Limit(host, hash string, limit int64) error {
	data := fmt.Sprintf("hashes=%s&limit=%d", hash, limit)
	_, err := http.Post(host+api_limit, "application/x-www-form-urlencoded; charset=UTF-8", strings.NewReader(data))
	return err
}

func Update_TK(host string) {
	resp, err := http.Get("https://raw.githubusercontent.com/ngosang/trackerslist/master/trackers_all.txt")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	jsondata := fmt.Sprintf(`{"add_trackers_enabled": true, "add_trackers": "%s"}`, url.QueryEscape(string(body)))

	data := fmt.Sprintf("json=%s", jsondata)
	resp, err = http.Post(host+api_setting, "application/x-www-form-urlencoded; charset=UTF-8", strings.NewReader(data))
	if resp.StatusCode == 200 && err == nil {
		log.Printf("[%s] trackers 更新成功!", host)
	}

}
