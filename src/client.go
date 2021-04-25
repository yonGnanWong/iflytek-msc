package src

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Client struct {
	c          *http.Client
	taskId     string
	path       string
	fileName   string
	fileLength string
}

func (c *Client) Prepare() {
	log.Println("prepare request start.....")

	ts := getTime()
	body := url.Values{}
	body.Add("app_id", AppId)
	body.Add("signa", getSig(ts))
	body.Add("ts", ts)
	body.Add("file_len", getFileSize(c.path))
	body.Add("file_name", c.fileName)
	body.Add("slice_num", "1")

	req := getRequest("POST", ApiPrepare, strings.NewReader(body.Encode()))
	result := c.DoReq(req)

	c.taskId = result.Data

	log.Printf("prepare request done, result info:%v , taskId:%s \n\n", result, result.Data)
}

func (c Client) Upload() {
	log.Println("upload request start.....")

	file, err := os.Open(c.path)
	assert1(err)

	ts := getTime()
	params := make(map[string]string)
	params["app_id"] = AppId
	params["signa"] = getSig(ts)
	params["ts"] = ts
	params["task_id"] = c.taskId
	params["slice_id"] = getNextSliceId()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileWriter, err := writer.CreateFormFile("content", "test.m4a")
	assert1(err)

	_, err = io.Copy(fileWriter, file)
	assert1(err)

	for key, val := range params {
		if err := writer.WriteField(key, val); err != nil {
			assert1(err)
		}
	}

	_ = writer.Close()
	req := getRequest("POST", ApiUpload, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	result := c.DoReq(req)

	log.Printf("upload request done, result info:%v \n\n", result)
}

func (c Client) Merge() {
	log.Println("merge request start.....")

	ts := getTime()
	body := url.Values{}
	body.Add("app_id", AppId)
	body.Add("signa", getSig(ts))
	body.Add("ts", ts)
	body.Add("task_id", c.taskId)

	req := getRequest("POST", ApiMerge, strings.NewReader(body.Encode()))
	result := c.DoReq(req)

	log.Printf("merge request done, result info:%v \n\n", result)
}

func (c Client) GetProgress() {
	log.Println("getProgress request start.....")
	ticker := time.NewTicker(1 * time.Second)

LOOP:
	for  {
		select {
		case <- ticker.C:
			ts := getTime()

			body := url.Values{}
			body.Add("app_id", AppId)
			body.Add("signa", getSig(ts))
			body.Add("ts", ts)
			body.Add("task_id", c.taskId)

			req := getRequest("POST", ApiGetProgress, strings.NewReader(body.Encode()))
			result := c.DoReq(req)

			if result.ErrNo != 0 && result.ErrNo != 26605 {
				log.Printf("getProgress request failed, code:%d,message:%v \n", result.ErrNo, result.Failed)
				break LOOP
			} else {

				status := new(status)
				json.Unmarshal([]byte(result.Data), &status)

				if status.Status == 9 {
					log.Printf("task is done,status:%v desc:%s", status.Status, status.Desc)
					break LOOP
				}

				log.Printf("task is running,status:%v desc:%s", status.Status, status.Desc)
			}
		}
	}
}

func (c Client) GetResult() {
	log.Println("getResult request start.....")

	ts := getTime()
	body := url.Values{}
	body.Add("app_id", AppId)
	body.Add("signa", getSig(ts))
	body.Add("ts", ts)
	body.Add("task_id", c.taskId)

	req := getRequest("POST", ApiGetResult, strings.NewReader(body.Encode()))
	result := c.DoReq(req)
	log.Printf("getResult request done, result info:%v \n\n", result)
}

func (c Client) DoReq(req *http.Request) *result {
	response, err := c.c.Do(req)
	assert1(err)
	byt, err := ioutil.ReadAll(response.Body)
	assert1(err)

	resp := new(result)
	_ = json.Unmarshal(byt, resp)

	return resp
}

/**
new client, 返回默认类型和传新的client
*/
func newClient() *Client {
	return &Client{
		c:        new(http.Client),
		path:     "./test.m4a",
		fileName: fileName,
	}
}
