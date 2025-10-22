package main

import (
	"bytes"
	"fmt"
	"io"

	"net/http"
	"os"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"

	"github.com/opentelekomcloud-community/otc-api-sign-sdk-go/core"
)

func ecsStart(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	s := core.Signer{
		Key:           ctx.GetSecurityAccessKey(),
		Secret:        ctx.GetSecuritySecretKey(),
		SecurityToken: ctx.GetSecurityToken(),
	}

	// get project ID from environment variables
	project_id := os.Getenv("RUNTIME_PROJECT_ID")

	var endpoint = ctx.GetUserData("ECS_ENDPOINT")
	if endpoint == "" {
		endpoint = "ecs.eu-de.otc.t-systems.com"
	}

	instance_id := ctx.GetUserData("INSTANCE_ID")

	var jsonStr = []byte(`{"os-start": {"servers": [{ "id": "` + instance_id + `" }] }}`)

	request, err := http.NewRequest("POST", "https://"+endpoint+"/v1/"+project_id+"/cloudservers/action",
		io.NopCloser(bytes.NewBuffer(jsonStr)))

	if err != nil {
		ctx.GetLogger().Logf("ECS start request failed:", err)
		return "invalid request", err
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

		ctx.GetLogger().Logf("ECS start request failed:", err)
		return "invalid request", err
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

	ret := map[string]interface{}{
		"statusCode":      resp.Status,
		"headers":         resp.Header,
		"isBase64Encoded": false,
		"body":            string(body),
	}

	ctx.GetLogger().Logf("Response:", ret)

	return ret, nil
}

func main() {
	runtime.Register(ecsStart)
}
