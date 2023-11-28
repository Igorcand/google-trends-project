package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/hhtp"
	"os"
)

type RSS struct{
	XMLName		xml.Name 		`xml: "rss"`
	Channel 	*Channel		`xml: "channel"`
}

type Channel struct{
	Title		string 			`xml: "title"`
	ItemList	[]Item			`xml: "item"`
}

type Item struct {
	Title 		string			`xml: "title"`
	Link 		string 			`xml: "link"`
	Traffic 	string  		`xml: "approx_traffic"`
	NewsItems 	[]News 			`xml: "new_item"`
}

type News struct{
	Headline 		string		`xml: "new_item_title"`
	Headlinelink 	string		`xml: "new_item_url"`
}

func main(){
	var r RSS 

	data := readGoogleTrends()
}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com.br/trends/trendingsearches/daily/rss?geo=US")

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}

func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}
