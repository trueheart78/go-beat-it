package search

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36"
	hltbURL   = "https://howlongtobeat.com/search_main.php?&page=1"
)

func Submit(name string, url string) (html io.Reader, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(formData(name)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://howlongtobeat.com")
	req.Header.Set("Referer", "https://howlongtobeat.com/")
	req.Header.Set("User-Agent", userAgent)

	res, _ := client.Do(req)

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Response is not OK (%d)", res.StatusCode)
	}
	html = req.Body
	return
}

func formData(name string) string {
	values := make(url.Values)

	values.Set("queryString", name)
	values.Set("t", "games")
	values.Set("sorthead", "popular")
	values.Set("sortd", "Normal Order")
	values.Set("plat", "")
	values.Set("length_type", "main")
	values.Set("length_min", "")
	values.Set("length_max", "")
	values.Set("detail", "")

	return values.Encode()
}
