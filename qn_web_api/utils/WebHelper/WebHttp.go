package WebHelper

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
	"qnsoft/qn_web_api/utils/ErrorHelper"

	"github.com/ajg/form"
)

/*
*post请求
 */
func HttpPost(_url string, _headers map[string]string, _body map[string]interface{}) string {
	var _rt = ""
	_formData, err := form.EncodeToString(_body)
	ErrorHelper.CheckErr(err)
	if len(_url) > 0 && len(_body) > 0 {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		payload := strings.NewReader(_formData)
		req, _ := http.NewRequest("POST", _url, payload)
		if len(_headers) <= 0 {
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded; text/html; charset=utf-8")
			req.Header.Add("Accept", "text/html")
		}
		for k, v := range _headers {
			req.Header.Add(k, v)
		}
		res, err := client.Do(req)
		ErrorHelper.CheckErr(err)
		defer res.Body.Close()
		if res != nil {
			body, _ := ioutil.ReadAll(res.Body)
			_rt = string(body)
		}
	}
	return _rt
}

/*
*get请求
 */
func HttpGet(_url string, _headers map[string]string, _body map[string]interface{}) string {
	var _rt = ""
	_formData, err := form.EncodeToString(_body)
	ErrorHelper.CheckErr(err)
	if len(_url) > 0 && len(_body) > 0 {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		payload := strings.NewReader(_formData)
		req, _ := http.NewRequest("GET", _url, payload)
		if len(_headers) <= 0 {
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded; text/html; charset=utf-8")
			req.Header.Add("Accept", "text/html")
		}
		for k, v := range _headers {
			req.Header.Add(k, v)
		}
		res, err := client.Do(req)
		ErrorHelper.CheckErr(err)
		defer res.Body.Close()
		if res != nil {
			body, _ := ioutil.ReadAll(res.Body)
			_rt = string(body)
		}
	}
	return _rt
}
