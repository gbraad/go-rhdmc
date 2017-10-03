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

	rhdmc "github.com/gbraad/go-rhdmc"
)

func main() {
    username := flag.String("username", "", "RHD Account username (required)")
    password := flag.String("password", "", "RHD Account password (required)")
    filepath := flag.String("filepath", "", "Donwload manager filepath (required)")
    flag.Parse()
    
    if *username == "" || *password == "" || *filepath == "" {
        fmt.Println("Missing required information\nUse -h for more information")
        os.Exit(1)
    }
    
	filename, err:= rhdmc.Download(*username, *password, *filepath)
	
	if err == nil {
		fmt.Printf("File has been download as %s\n", filename)
	
	    os.Exit(0)
	} else {
		fmt.Printf("Error occured: %s\n", err.Error())
		
		os.Exit(2)
	}
}