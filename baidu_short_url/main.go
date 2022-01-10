package baiduShortURL

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// ShortURLRequest 短链请求参数
type ShortURLRequest struct {
	LongURL        string         `json:"LongUrl"`
	TermOfValidity TermOfValidity `json:"TermOfValidity"`
}

// ShortURLData 短链返回详情
type ShortURLData struct {
	Code     int64  `json:"Code"`
	ErrMsg   string `json:"ErrMsg"`
	LongURL  string `json:"LongUrl"`
	ShortURL string `json:"ShortUrl"`
}

// ShortURLHttpResponse 短链返回
type ShortURLHttpResponse struct {
	Code   int64          `json:"Code"`
	ErrMsg string         `json:"ErrMsg"`
	Detail []ShortURLData `json:"ShortUrls"`
}

// Baidu 百度实例
type Baidu struct {
	Token string
}

// TermOfValidity 有效期
type TermOfValidity string

const (
	// ValidityLongTerm 长期
	ValidityLongTerm TermOfValidity = `long-term`
	// ValidityOneYear 一年
	ValidityOneYear TermOfValidity = `1-year`
)

// NewBaidu 实例化百度
func NewBaidu() *Baidu {
	return &Baidu{Token: `xxxxx`}
}

// ShortURLResponse 短链返回
type ShortURLResponse struct {
	URL          string                `json:"url"`
	HTTPResponse *ShortURLHttpResponse `json:"http_response"`
}

// ShortURL 短链地址
func (b *Baidu) ShortURL(url string, validity TermOfValidity) (*ShortURLResponse, error) {
	params := make([]ShortURLRequest, 0)
	params = append(params, ShortURLRequest{
		LongURL:        url,
		TermOfValidity: validity,
	})

	jsonStr, jsonErr := json.Marshal(params)
	if jsonErr != nil {
		return nil, jsonErr
	}

	request, requestErr := http.NewRequest("POST", `https://dwz.cn/api/v3/short-urls`, bytes.NewBuffer(jsonStr))
	if requestErr != nil {
		return nil, requestErr
	}

	request.Header.Add("Content-Type", `application/json; charset=UTF-8`)
	request.Header.Add("Dwz-Token", b.Token)
	defer request.Body.Close()

	client := &http.Client{Timeout: 3 * time.Second}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return nil, responseErr
	}
	defer response.Body.Close()

	content, contentErr := ioutil.ReadAll(response.Body)
	if contentErr != nil {
		return nil, contentErr
	}

	var result ShortURLHttpResponse
	if err := json.Unmarshal(content, &result); nil != err {
		return nil, err
	}

	res := &ShortURLResponse{
		URL:          url,
		HTTPResponse: &result,
	}

	if result.Code == 0 {
		for _, v := range result.Detail {
			if v.Code == 0 && v.LongURL == url {
				res.URL = v.ShortURL
			}
		}
	}

	return res, nil
}
