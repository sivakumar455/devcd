package utils

import (
	"bytes"
	"devcd/logger"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

func CreateCbBucket(cbHost, bucketName, bucketType string) (*http.Response, error) {

	data := url.Values{}
	data.Set("name", bucketName)
	data.Set("ramQuota", "110")
	data.Set("bucketType", bucketType)
	data.Set("flushEnabled", "1")
	data.Set("replicaNumber", "0")

	// curl request to go code curl -v -X POST http://10.143.201.101:8091/pools/default/buckets
	// 	-u Administrator:password
	// 	-d name=testBucket
	// 	-d ramQuota=256
	// 	-d bucketType=ephemeral

	cbUrl := fmt.Sprintf("%s/pools/default/buckets", cbHost)
	resp, err := PostAppUrlEncode(cbUrl, data, true)

	if err != nil {
		logger.Error("Error reading response body ", "error", err)
		return nil, err
	}

	return resp, nil
}

func CreateCbScope(cbHost, bucketName, scopeName string) (*http.Response, error) {

	data := url.Values{}
	data.Set("name", scopeName)

	// curl -X POST -v -u Administrator:password \
	// 	http://10.143.210.101:8091/pools/default/buckets/testBucket/scopes \
	// 	-d name=my_scope

	cbUrl := fmt.Sprintf("%s/pools/default/buckets/%s/scopes", cbHost, bucketName)
	resp, err := PostAppUrlEncode(cbUrl, data, true)

	if err != nil {
		logger.Error("Error reading response body ", "error", err)
		return nil, err
	}

	return resp, nil

}

func CreateCbCollection(cbHost, bucketName, scopeName, collectionName string) (*http.Response, error) {

	data := url.Values{}
	data.Set("name", collectionName)

	// curl -X POST -v -u Administrator:password \
	// 	http://10.143.210.101:8091/pools/default/buckets/\
	// 	testBucket/scopes/my_scope/collections \
	// 	-d name=my_collection_in_my_scope \
	// 	-d maxTTL=63113904 \
	// 	-d history=false

	cbUrl := fmt.Sprintf("%s/pools/default/buckets/%s/scopes/%s/collections", cbHost, bucketName, scopeName)
	resp, err := PostAppUrlEncode(cbUrl, data, true)

	if err != nil {
		logger.Error("Error reading response body:", "error", err)
		return nil, err
	}

	return resp, nil

}

func CreateCbIndex(cbHost, indexStr string) (*http.Response, error) {

	data := url.Values{}
	data.Set("statement", indexStr)

	cbUrl := fmt.Sprintf("%s/query/service", cbHost)
	logger.Debug("Index Query: ", "query", data)
	resp, err := PostAppUrlEncode(cbUrl, data, true)

	if err != nil {
		logger.Error("Error reading response body:", "error", err)
		return nil, err
	}

	return resp, nil

}

func InsertCbDoc(cbHost, cbDocStr string) (*http.Response, error) {

	data := url.Values{}
	data.Set("statement", cbDocStr)

	cbUrl := fmt.Sprintf("%s/query/service", cbHost)
	resp, err := PostAppUrlEncode(cbUrl, data, true)

	if err != nil {
		logger.Error("Error reading response body:", "error", err)
		return nil, err
	}

	return resp, nil

}

func PostAppUrlEncode(cbUrl string, data url.Values, auth bool) (*http.Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", cbUrl, strings.NewReader(data.Encode()))
	if err != nil {
		logger.Error("Error creating request:", "error", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if auth {
		req.SetBasicAuth("Administrator", "Administrator")
	}

	// reqDump, err := httputil.DumpRequestOut(req, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("\nREQUEST:\n%s\n", string(reqDump))
	logger.Debug("Request for index:", "request", req)
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error sending request to server:", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	logger.Info(fmt.Sprintf("Response status code: %v", resp.StatusCode))

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return nil, err
	// }

	// fmt.Println("Response body:", string(body))

	return resp, nil

}

func createBucketWithJsonBody(bucketName string, bucketType string) (*http.Response, error) {
	logger.Info(fmt.Sprintf("Creating %s bucket %s\n", bucketType, bucketName))
	data := map[string]interface{}{
		"ramQuotaMB":    110,
		"flushEnabled":  1,
		"bucketType":    bucketType,
		"replicaNumber": 0,
		"name":          bucketName,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Error("Error marshaling JSON data:", "error", err)
		return nil, err
	}

	logger.Debug(fmt.Sprintf("jsonData:%s", jsonData))

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8091/pools/default/buckets", bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Error("Error creating request:", "error", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("Administrator", "Administrator")

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error sending request to server:", "error", err)
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body:", "error", err)
		return nil, err
	}

	logger.Info(fmt.Sprintf("Response status code: %v", resp.StatusCode))
	logger.Info(fmt.Sprintf("Response body %v", string(bodyBytes)))

	return resp, nil
}

func RetryHttp(attempts int, sleep time.Duration, f func() (*http.Response, error)) error {
	var err error
	var resp *http.Response
	for i := 0; i < attempts; i++ {
		resp, err = f()
		if err == nil {
			return nil
		}
		logger.Info("Retrying after error ", "error", err)
		time.Sleep(sleep)
		sleep *= 2
	}
	return fmt.Errorf("after %d attempts, resp : %v last error: %v", attempts, resp, err)
}

func CreateCBContainer(containerRtEngine, containerName, volumeName string) {

	// Create and start the container
	cmd := exec.Command(containerRtEngine, "run", "-d", "--name", containerName,
		"-p", "8091-8096:8091-8096", "-p", "11210-11211:11210-11211",
		"-v", fmt.Sprintf("%s:/opt/couchbase/var", volumeName), "couchbase")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		logger.Error("Error creating Couchbase container", "error", err)
		return
	}
}
