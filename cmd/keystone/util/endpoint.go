package util

// TODO replace to correct folader

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/itouri/sgx-iaas/pkg/domain/keystone"
)

var (
	endpoints        map[keystone.EnumServiceType]string
	keystoneEndpoint string
)

func GetEndPoint(est keystone.EnumServiceType) (string, error) {
	if endpoints == nil {
		endpoints = map[keystone.EnumServiceType]string{}
	}

	if ep, ok := endpoints[est]; ok {
		return ep, nil
	}

	// TODO 非ハードコーディング
	if keystoneEndpoint == "" {
		keystoneEndpoint = GetEndpointURL()
	}

	// TODO clientの通信の部分を作り込む
	return ResolveServiceEndpoint("http://"+keystoneEndpoint, est)
}

// TODO 非ハードコーディング
func GetEndpointURL() string {
	return "localhost:1323"
}

func RegisterEndpoint(t keystone.EnumServiceType, ipaddr string, port uint64) error {
	service := &keystone.Service{
		Type:   t,
		Port:   port,
		IPAddr: ipaddr, // エラー返さないの？
	}

	sJSON, err := json.Marshal(service)
	if err != nil {
		return err
	}

	// debug
	fmt.Println(string(sJSON))

	endpoint := GetEndpointURL()
	resp, err := http.Post("http://"+endpoint+"/v1/services", "application/json", bytes.NewBuffer(sJSON))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Posting falied: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}

type Req struct {
	/// Name   string `json:"name"`
	Type   int    `json:"type"`
	Port   int    `json:"port"`
	IPAddr string `json:"ipaddr"`
}

func ResolveServiceEndpoint(endpointURL string, st keystone.EnumServiceType) (string, error) {
	// heatに情報を送るためにはendpointからIPを解決する必要がある
	resp, err := http.Get(endpointURL + "/v1/services/resolve/" + st.String())
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code is %d", resp.StatusCode)
	}

	req := new(Req)
	err = decodeJSON(resp, req)
	if err != nil {
		return "", err
	}

	url := "http://" + req.IPAddr + ":" + strconv.Itoa(req.Port)
	return url, nil
}

func decodeJSON(resp *http.Response, v interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// TODO なぜBASE64エンコードされている？
	b := strings.Replace(string(body), "\"", "", -1)
	str, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(str), v)
	if err != nil {
		return err
	}

	return nil
}
