/*
Copyright (C) 2017 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rhdmc

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	DMURL = "https://developers.redhat.com/download-manager/jdf/file/"
)

func Download(username string, password string, filename string) (bool, error) {
	var url string = fmt.Sprintf("%s%s?workflow=direct", DMURL, filename)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("Error in request: '%s'", err.Error())
	}
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("Unable to get resource '%s': %s", filename, err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("Wrong credentials or filename")
	}

	defer func() { _ = resp.Body.Close() }()
	out, err := os.Create(filename)
	defer out.Close()
	if err != nil {
		return false, fmt.Errorf("Not able to create file as '%s': %s", filename, err.Error())
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return false, fmt.Errorf("Not able to copy file to '%s': %s", filename, err.Error())
	}

	return true, nil
}
