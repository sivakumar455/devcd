package couchbase

import (
	"devcd/logger"
	"devcd/utils"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func CouchbaseInit() {
	CbClusterInit()
	CbBucketInit()
	CbIndexInit()
}

func CbClusterInit() {

	logger.Info("Creating cluster", "cluster", CBConfig.ClusterInfo.ClusterName)

	baseUrl := CBConfig.ClusterInfo.BaseURL
	cbServices := CBConfig.ClusterInfo.Services
	memoryQuota := fmt.Sprintf("%v", CBConfig.ClusterInfo.MemoryQuota)
	indexMemoryQuota := fmt.Sprintf("%v", CBConfig.ClusterInfo.IndexMemoryQuota)
	clusterName := CBConfig.ClusterInfo.ClusterName
	storageMode := CBConfig.ClusterInfo.StorageMode
	cbUsername := CBConfig.ClusterInfo.Username
	cbPassword := CBConfig.ClusterInfo.Password
	cbPort := CBConfig.ClusterInfo.Port

	// curl  -v -X POST http://127.0.0.1:8091/node/controller/setupServices -d 'services=kv%2Cn1ql%2Cindex'
	// curl  -v -X POST http://127.0.0.1:8091/pools/default -d 'memoryQuota=256' -d 'indexMemoryQuota=256'
	// curl  -u Administrator:Administrator -v -X POST http://127.0.0.1:8091/settings/web -d 'password=Administrator&username=Administrator&port=SAME'

	// First request
	postUrl := fmt.Sprintf("%s/node/controller/setupServices", baseUrl)
	data := url.Values{}
	data.Set("services", cbServices)

	err := utils.RetryHttp(5, 5*time.Second, func() (*http.Response, error) {
		return utils.PostAppUrlEncode(postUrl, data, false)
	})
	if err != nil {
		logger.Error("Error:", "error", err)
	}

	// Second request
	url2 := fmt.Sprintf("%s/pools/default", baseUrl)
	data2 := url.Values{}
	data2.Set("memoryQuota", memoryQuota)
	data2.Set("indexMemoryQuota", indexMemoryQuota)
	utils.PostAppUrlEncode(url2, data2, false)

	// Third request
	url3 := fmt.Sprintf("%s/settings/web", baseUrl)
	data3 := url.Values{}
	data3.Set("username", cbUsername)
	data3.Set("password", cbPassword)
	data3.Set("port", cbPort)
	utils.PostAppUrlEncode(url3, data3, true)

	// Fourth request
	url4 := fmt.Sprintf("%s/pools/default", baseUrl)
	data4 := url.Values{}
	data4.Set("clusterName", clusterName)
	utils.PostAppUrlEncode(url4, data4, true)

	// Fifth request
	url5 := fmt.Sprintf("%s/settings/indexes", baseUrl)
	data5 := url.Values{}
	data5.Set("storageMode", storageMode)
	utils.PostAppUrlEncode(url5, data5, true)

	logger.Info("Created cluster ", "cluster", CBConfig.ClusterInfo.ClusterName)

}

func CbBucketInit() {
	logger.Info("Creating Buckets,Scopes and Collections", "cluster", CBConfig.ClusterInfo.ClusterName)
	baseUrl := CBConfig.ClusterInfo.BaseURL

	for _, bucket := range CBConfig.BucketInfo {

		logger.Info("Creating bucket ", "bucket", bucket.Name)
		utils.CreateCbBucket(baseUrl, bucket.Name, bucket.Type)

		for _, scope := range bucket.Scope {
			logger.Info("Creating Scope ", "scope", scope.Name)
			utils.CreateCbScope(baseUrl, bucket.Name, scope.Name)

			for _, collection := range scope.Collection {
				logger.Info("Creating Collection ", "collection", collection.Name)
				utils.CreateCbCollection(baseUrl, bucket.Name, scope.Name, collection.Name)
			}
		}
	}
	logger.Info("Created Buckets,Scopes and Collections", "cluster", CBConfig.ClusterInfo.ClusterName)
}

func CbIndexInit() {
	logger.Info("Creating indexes", "cluster", CBConfig.ClusterInfo.ClusterName)
	baseUrl := "http://127.0.0.1:8093"

	for _, bucket := range CBConfig.BucketInfo {
		for _, scope := range bucket.Scope {
			for _, collection := range scope.Collection {
				logger.Info(fmt.Sprintf("Creating index for : %s.%s.%s ", bucket.Name, scope.Name, collection.Name))
				for _, index := range collection.Index {
					utils.CreateCbIndex(baseUrl, index)
				}
			}
		}
	}
	logger.Info("Created indexes", "cluster", CBConfig.ClusterInfo.ClusterName)
}
