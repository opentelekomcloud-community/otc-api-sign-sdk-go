/* Sample code to start an ECS instance using AK/SK authentication and the Open Telekom Cloud API Signing SDK for Go.
*
* Before running this example, please set the environment variables
* OTC_SDK_AK, OTC_SDK_SK, OTC_SDK_PROJECT_ID and ECS_INSTANCE_ID in the local environment.
*
* - OTC_SDK_AK: Your Access Key
* - OTC_SDK_SK: Your Secret Key
* - OTC_SDK_PROJECT_ID: Your Project ID
* - ECS_INSTANCE_ID: ID of the ecs instance to start
 *
*/
package main

import (
	"bytes"
	"fmt"
	"io"

	"net/http"
	"os"

	"github.com/opentelekomcloud-community/otc-api-sign-sdk-go/core"
)

func main() {
	demoStartECS()
}

func demoStartECS() {
	//If ak and sk used for authentication are hard-coded into the code or stored in plain text,
	// there is a great security risk. It is recommended to store the cipher text in the configuration
	// file or environment variable, and decrypt it when used to ensure security;
	//This example uses ak and sk stored in environment variables.
	// Before running this example, please set the environment variables
	// OTC_SDK_AK and OTC_SDK_SK in the local environment.
	s := core.Signer{
		Key:    os.Getenv("OTC_SDK_AK"),
		Secret: os.Getenv("OTC_SDK_SK"),
	}

	instance_id := os.Getenv("ECS_INSTANCE_ID")

	project_id := os.Getenv("OTC_SDK_PROJECT_ID")

	var jsonStr = []byte(`{"os-start": {"servers": [{ "id": "` + instance_id + `" }] }}`)

	request, err := http.NewRequest("POST", "https://ecs.eu-de.otc.t-systems.com/v1/"+project_id+"/cloudservers/action",
		io.NopCloser(bytes.NewBuffer(jsonStr)))

	if err != nil {
		fmt.Println(err)
		return
	}

	request.Header.Set("content-type", "application/json; charset=utf-8")

	if project_id != "" {
		// To access resources in a sub-project (e.g. eu_de/myproject)
		// by calling APIs, X-Project-Id of "eu_de/myproject" is needed
		request.Header.Set("X-Project-Id", project_id)
	}

	// Sign the request
	s.Sign(request)

	client := http.DefaultClient

	// Send the request
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", string(body))
}
