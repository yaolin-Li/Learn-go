package main

import (
	"get-bli/downloader"
	myfmt "get-bli/fmt"
)

func main() {
	request := downloader.InfoRequest{Bvids: []string{"BV1zK4y1978Z","BV1Ff4y187q9"}}
	response, err := downloader.BatchDownLoadVideoInfo(request); if err != nil {
		panic(err)
	}
	for _, info := range response.Infos {
		myfmt.Logger.Printf("title: %s \ndesc: %s\n",info.Data.Title,info.Data.Desc)
	}
}