package util

// TODO replace to correct folader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

	// TODO clientの通信の部分を作り込む
	return ResolveServiceEndpoint(keystoneEndpoint, est)
}

// TODO 非ハードコーディング
func GetEndpointURL() string {
	return "192.168.0.2:1323"
}

func ResolveServiceEndpoint(endpointURL string, st keystone.EnumServiceType) (string, error) {
	// heatに情報を送るためにはendpointからIPを解決する必要がある
	resp, err := http.Get(endpointURL + "/services/resolve/" + st.String())
	if err != nil {
		return "", err
	}

	service := &keystone.Service{}
	err = decodeJSON(resp, service)
	if err != nil {
		return "", err
	}

	url := "http://" + service.IPAddr.String() + ":" + string(service.Port)
	return url, nil
}

func decodeJSON(resp *http.Response, v interface{}) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
