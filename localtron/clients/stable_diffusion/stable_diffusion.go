/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package stable_diffusion

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Address string
}

func NewClient(address string) *Client {
	return &Client{Address: address}
}

// represents
// "data":["draw me a cat",1,6,256,256,7.5,0,false,false,"PNDM",0.25,null,null,null]
type StableDiffusionParams struct {
	Prompt        string  `json:"prompt"`
	NumImages     int     `json:"num_images"`
	Steps         int     `json:"steps"`
	Width         int     `json:"width"`
	Height        int     `json:"height"`
	GuidanceScale float64 `json:"guidance_scale"`
	Seed          int     `json:"seed"`
	Flag1         bool    `json:"flag1"`
	Flag2         bool    `json:"flag2"`
	Scheduler     string  `json:"scheduler"`
	Rate          float64 `json:"rate"`
	Optional1     *string `json:"optional1"`
	Optional2     *string `json:"optional2"`
	Optional3     *string `json:"optional3"`
}

type PredictRequest struct {
	FnIndex int `json:"fn_index"`
	/* Params gets turned into `Data` */
	Params      StableDiffusionParams `json:"-"`
	Data        []interface{}         `json:"data"`
	SessionHash string                `json:"session_hash"`
}

func (pr *PredictRequest) ConvertParamsToData() {
	pr.Data = []interface{}{
		pr.Params.Prompt,
		pr.Params.NumImages,
		pr.Params.Steps,
		pr.Params.Width,
		pr.Params.Height,
		pr.Params.GuidanceScale,
		pr.Params.Seed,
		pr.Params.Flag1,
		pr.Params.Flag2,
		pr.Params.Scheduler,
		pr.Params.Rate,
		pr.Params.Optional1,
		pr.Params.Optional2,
		pr.Params.Optional3,
	}
}

type FileData struct {
	Name   string      `json:"name"`
	Data   interface{} `json:"data"`
	IsFile bool        `json:"is_file"`
}

type HistoryData struct {
	Headers []string   `json:"headers"`
	Data    [][]string `json:"data"`
}

type PredictResponse struct {
	Data            [][]FileData `json:"data"`
	HistoryData     HistoryData  `json:"history_data"`
	IsGenerating    bool         `json:"is_generating"`
	Duration        float64      `json:"duration"`
	AverageDuration float64      `json:"average_duration"`
}

/*
FileURL returns the local stable diffusion URL of a fileName acquired through the stable diffusion API.

eg. http://127.0.0.1:8001/file=/tmp/tmppjk80zb6/tmpv4bk66e9.png
*/
func FileURL(addr string, fileName string) string {
	return fmt.Sprintf("%v:/file=%v", addr, fileName)
}

// GetImageAsBase64 fetches the image from the given URL and returns it as a base64 encoded string.
func GetImageAsBase64(imageURL string) (string, error) {
	resp, err := http.Get(imageURL)
	if err != nil {
		return "", errors.New("failed to fetch image: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch image: status code " + resp.Status)
	}

	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("failed to read image data: " + err.Error())
	}

	base64Image := base64.StdEncoding.EncodeToString(imageData)

	return base64Image, nil
}

func (c *Client) Predict(req PredictRequest) (*PredictResponse, error) {
	url := c.Address + "/run/predict/"

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Accept", "*/*")
	httpReq.Header.Set("Connection", "keep-alive")
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("API request failed with status code: " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var predictResp PredictResponse
	err = json.Unmarshal(body, &predictResp)
	if err != nil {
		return nil, err
	}

	return &predictResp, nil
}
