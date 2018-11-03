package main

import (
	"fmt"
	ss "github.com/shadowsocks/shadowsocks-go/shadowsocks"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

const Host = ""
const Port = 0
const Method = "aes-256-cfb"
const Password = ""

func useShadowSocks(client *http.Client, parsedURL *url.URL) *http.Client {

	address := fmt.Sprintf("%s:%d", Host, Port)

	host, _, err := net.SplitHostPort(parsedURL.Host)
	if err != nil {
		if parsedURL.Scheme == "https" {
			host = net.JoinHostPort(parsedURL.Host, "443")
		} else {
			host = net.JoinHostPort(parsedURL.Host, "80")
		}
	} else {
		host = parsedURL.Host
	}
	rawAddr, err := ss.RawAddr(host)

	cipher, err := ss.NewCipher(Method, Password)

	tr := &http.Transport{
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return ss.DialWithRawAddr(rawAddr, address, cipher)
	}

	client.Transport = tr
	return client
}

func main() {

	urlAddress := "https://twitter.com/?lang=ja"

	client := http.DefaultClient
	parsedURL, _ := url.Parse(urlAddress)
	client = useShadowSocks(http.DefaultClient, parsedURL)
	request, _ := http.NewRequest(http.MethodGet, urlAddress, nil)

	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error Do: %s", err)
		return
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error Read: %s", err)
		return
	}

	fmt.Println(string(content))

	log.Printf("StatusCode: %d\n", response.StatusCode)
}
