/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
)

// appsCmd represents the apps command
var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "List all applications inside Rapid Deploy",
	Long:  `List all applications inside Rapid Deploy`,
	Run: func(cmd *cobra.Command, args []string) {
		getApps()
	},
}

func init() {
	rootCmd.AddCommand(appsCmd)
}

type ApplicationResponse struct {
	ApplicationName string `json:"applicationName"`
	Guid            string `json:"guid"`
}

func getApps() {
	fmt.Println("Listing all apps...")
	url := "http://localhost:9000/api/application"
	responseBytes := getAppData(url)
	var app []ApplicationResponse
	if err := json.Unmarshal(responseBytes, &app); err != nil {
		errors := ansi.Color("Could not unmarshal response", "green+h:black")
		fmt.Println(string(errors))
	}

	for i := 0; i < len(app); i++ {
		fmt.Printf("App Name -> %s, Guid -> %s\n", app[i].ApplicationName, app[i].Guid)
	}
	msg := ansi.Color("OK", "green+b")
	fmt.Println(string(msg))
}

func getAppData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		errmsg := ansi.Color("Could not make request", "red+h")
		log.Printf(errmsg)
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		errmsg := ansi.Color("Could not unmarshal response", "red+h")
		log.Printf(errmsg)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	return responseBytes
}
