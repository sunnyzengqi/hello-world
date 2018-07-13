package main

import (
	"log"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"github.com/robfig/cron"
)

func main() {
	welcome()
	c:= cron.New()
	eatSpec := "0 0 12,18 * * *"//每天饭点执行
	meetingSpec := "0 0 11 * * 1"//每周星期1执行
	reportSpec := "0 0 16 * * 4"//每周星期4执行
	c.AddFunc(eatSpec,eat)
	c.AddFunc(meetingSpec,meeting)
	c.AddFunc(reportSpec,report)
	c.Start()
	select{}
}

type msg struct {
	Msgtype string  `json:"msgtype"`
	Text content  `json:"text"`
}

type content struct{
	Content string  `json:"content"`
}

var url = "https://oapi.dingtalk.com/robot/send?access_token=2b589ffbcecc1eb0ea60056af076f534157bde41604dae0cc13f5a656db2d05c"

var index = 1
var turnMap = [7]string{"双渤","逸轩","魏超","刘冰","曾奇","朝晖","晨涛"}

var textPrefix = "大家好，我又来啦! "

func welcome(){
	text := "hello all，本机器人上线啦，以后请多多关照"
	msg := msg{Msgtype:"text",Text:content{Content:text}}
	theMsg, _ := json.Marshal(msg)
	reqBody := []byte(string(theMsg))
	resp := doHttp("POST",url,reqBody)
	resBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("welcome response Body:", string(resBody))
}

//每周一定会议室
func meeting(){
	if(index>=len(turnMap)){
		index=0
	}
	text := textPrefix+"本周该"+turnMap[index]+"负责预定周五会议室和会邀呦，不要忘啦"
	index++
	msg := msg{Msgtype:"text",Text:content{Content:text}}
	theMsg, _ := json.Marshal(msg)
	reqBody := []byte(string(theMsg))
	resp := doHttp("POST",url,reqBody)
	resBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("meeting response Body:", string(resBody))
}

//每周四发发周报提醒
func report(){
	text := textPrefix+"提醒下，发周报啦，说完我就溜了溜了"
	msg := msg{Msgtype:"text",Text:content{Content:text}}
	theMsg, _ := json.Marshal(msg)
	reqBody := []byte(string(theMsg))
	resp := doHttp("POST",url,reqBody)
	resBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("report response Body:", string(resBody))
}

//每天饭点吃饭
func eat(){
	text := textPrefix+"本机器人饿啦，去吃饭吧"
	msg := msg{Msgtype:"text",Text:content{Content:text}}
	theMsg, _ := json.Marshal(msg)
	reqBody := []byte(string(theMsg))
	resp := doHttp("POST",url,reqBody)
	resBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("eat response Body:", string(resBody))
}

func doHttp(method string,url string,body []byte) *http.Response{
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp
}