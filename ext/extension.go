package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Extension struct {
	ExtensionName     string             `json:"extensionName"`
	Description       string             `json:"extensionDesc"`
	Versions          []ExtensionVersion `json:"versions"`
	Statistics        map[string]int     `json:"statistics"`
	PublisherUsername string             `json:"publisher.publisherName"`
}

type ExtensionVersion struct {
	Version             string `json:"version"`
	ArtifactDownloadURL string `json:"artifactDownloadUrl"`
}

func getVSCodeExtensions(maxPage int, pageSize int, session *http.Client) <-chan Extension {
	ch := make(chan Extension)

	go func() {
		defer close(ch)

		for page := 1; page <= maxPage; page++ {
			body := map[string]interface{}{
				"filters": []map[string]interface{}{
					{
						"criteria": []map[string]interface{}{
							{
								"filterType": 8,
								"value":      "Microsoft.VisualStudio.Code",
							},
						},
						"pageNumber": page,
						"pageSize":   pageSize,
						"sortBy":     0,
						"sortOrder":  0,
					},
				},
				"assetTypes": []interface{}{},
				"flags":      0x1 | 0x2 | 0x4 | 0x8 | 0x10 | 0x20 | 0x40 | 0x80 | 0x100 | 0x200 | 0x1000 | 0x8000,
			}

			req, err := http.NewRequest("POST", "https://marketplace.visualstudio.com/_apis/public/gallery/extensionquery", nil)
			if err != nil {
				fmt.Printf("Error creating request: %v\n", err)
				return
			}

			req.Header.Set("Accept", "application/json; charset=utf-8; api-version=7.2-preview.1")
			req.Header.Set("Content-Type", "application/json")

			resp, err := session.Do(req.WithContext(req.Context()))
			if err != nil {
				fmt.Printf("Error making request: %v\n", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Printf("Error response: %d\n", resp.StatusCode)
				return
			}

			var response struct {
				Results []struct {
					Extensions []Extension `json:"extensions"`
				} `json:"results"`
			}

			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				fmt.Printf("Error decoding response: %v\n", err)
				return
			}

			for _, ext := range response.Results[0].Extensions {
				ch <- ext
			}

			if len(response.Results[0].Extensions) < pageSize {
				break
			}
		}
	}()

	return ch
}

func extension() {
	retryStrategy := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 10,
			MaxConnsPerHost:     10,
		},
	}

	for ext := range getVSCodeExtensions(10, 10, retryStrategy) {
		fmt.Printf("Extension Name: %s\n", ext.ExtensionName)
		fmt.Printf("Description: %s\n", ext.Description)
		for _, v := range ext.Versions {
			fmt.Printf("Version: %s, Download URL: %s\n", v.Version, v.ArtifactDownloadURL)
		}
		fmt.Printf("Publisher Username: %s\n", ext.PublisherUsername)
		fmt.Printf("Install Count: %d\n", ext.Statistics["install"])
		fmt.Println()
	}
}
