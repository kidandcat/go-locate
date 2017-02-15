package locate

import (
	"io/ioutil"
	"net/http"
)

func Locate(ip string) (string, string) {
	url := "http://ip-api.com/json/" + ip
	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), resp.Status
}
