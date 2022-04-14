package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

var port string
var myip bool

func main() {
	flag.BoolVar(&myip, "ip", false, "get my ip")
	flag.StringVar(&port, "port", "", "check my port")
	flag.Parse()
	if myip && port == "" {
		fmt.Println(getIP())
		os.Exit(0)
	} else if port != "" && !myip {
		ip := getIP()
		data := url.Values{
			"IP":   {ip},
			"port": {port},
		}
		resp, err := http.PostForm("https://canyouseeme.org/", data)

		if err != nil {
			log.Fatal(err)
		}
		bd, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		strbd := string(bd)
		if strings.Contains(strbd, "Error:") {
			fmt.Println("Close")
		} else if strings.Contains(strbd, "Success") {
			fmt.Println("Open")
		}
	} else {
		fmt.Println("please select only on ip or port")
	}

}

func getIP() string {
	var ip string
	c := colly.NewCollector(
		colly.AllowedDomains("canyouseeme.org"),
	)
	c.OnHTML("#ip", func(h *colly.HTMLElement) {
		ip = h.Attr("value")
	})
	c.Visit("https://canyouseeme.org")
	return ip
}
