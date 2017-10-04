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

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	rhdmc "github.com/gbraad/go-rhdmc"
)

func main() {
	var versions = []struct {
		value    string
		released string
	}{
		{value: "3.6.173.0.21", released: "2017-09-08"},
		{value: "3.5.5.31.24", released: "2017-09-07"},
	}

	fmt.Println("OpenShift Client Tools downloader\n")

	username := flag.String("username", "", "RHD Account username (required)")
	password := flag.String("password", "", "RHD Account password (required)")
	version := flag.String("version", versions[0].value, "Version to download")
	platformFlag := flag.String("platform", runtime.GOOS, "Platform (linux, windows, macosx")
	list := flag.Bool("list", false, "List available versions")
	flag.Parse()

	if *list {
		fmt.Println("Suggested versions")
		fmt.Println("--   version   ---  released  --")
		for _, v := range versions {
			fmt.Printf("%s\t| %s \n", v.value, v.released)
		}
		os.Exit(0)
	}

	// required
	if *username == "" || *password == "" {
		fmt.Println("Missing required information\nUse -h for more information")
		os.Exit(1)
	}

	// check platform
	platform := strings.ToLower(*platformFlag)
	if platform == "darwin" {
		platform = "macosx"
	}
	if platform != "linux" && platform != "windows" && platform != "macosx" {
		fmt.Println("Incorrect platform chosen\nUse -h for more information")
		os.Exit(1)
	}

	extension := "tar.gz"
	if platform == "windows" {
		extension = "zip"
	}

	// oc-3.5.5.31.24-windows.zip
	filepath := fmt.Sprintf("oc-%s-%s.%s", *version, platform, extension)

	fmt.Printf("Downloading '%s' ... ", filepath)

	_, err := rhdmc.Download(*username, *password, filepath)

	if err == nil {
		fmt.Println("OK")

		os.Exit(0)
	} else {
		fmt.Printf("FAIL\n%s", err.Error())

		os.Exit(2)
	}
}
