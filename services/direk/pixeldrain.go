package direk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/imroc/req/v3"
)

func formatSizePixeldrain(amount float64, precision int) string {
	if precision < 3 {
		precision = 3
	}
	switch {
	case amount >= 1e18:
		return fmt.Sprintf("%."+strconv.Itoa(precision)+"f EB", amount/1e18)
	case amount >= 1e15:
		return fmt.Sprintf("%."+strconv.Itoa(precision)+"f PB", amount/1e15)
	case amount >= 1e12:
		return fmt.Sprintf("%."+strconv.Itoa(precision)+"f TB", amount/1e12)
	case amount >= 1e9:
		return fmt.Sprintf("%."+strconv.Itoa(precision-1)+"f GB", amount/1e9)
	case amount >= 1e6:
		return fmt.Sprintf("%."+strconv.Itoa(precision-1)+"f MB", amount/1e6)
	case amount >= 1e3:
		return fmt.Sprintf("%."+strconv.Itoa(precision-1)+"f kB", amount/1e3)
	default:
		return fmt.Sprintf("%."+strconv.Itoa(precision)+"f B", amount)
	}
}

func Pixeldrain(url string) (string, string, error) {

	// checking the url
	url = strings.TrimPrefix(url, "/")
	url = strings.TrimSuffix(url, "/")

	parts := strings.Split(url, "/")
	fileID := parts[len(parts)-1]

	var infoLink, dlLink, filename string
	if parts[len(parts)-2] == "l" {
		infoLink = fmt.Sprintf("https://pixeldrain.com/api/list/%s", fileID)
		dlLink = fmt.Sprintf("https://pixeldrain.com/api/list/%s/zip?download", fileID)
	} else {
		infoLink = fmt.Sprintf("https://pixeldrain.com/api/file/%s/info", fileID)
		dlLink = fmt.Sprintf("https://pixeldrain.com/api/file/%s?download", fileID)
	}

	client := req.C() // Create a new client

	resp, err := client.R().Get(infoLink)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", "", err
	}
	var bandwidthUsedPaidFloat64 float64
	if bandwidthUsedPaid, ok := response["bandwidth_used_paid"].(float64); ok {
		bandwidthUsedPaidFloat64 = bandwidthUsedPaid
	} else {
		return "", "", errors.New("failed to parse bandwidth_used_paid from response")
	}

	var bandwidthLeftStr string
	bandwidthUsedPaid := int64(bandwidthUsedPaidFloat64)
	if bandwidthUsedPaid < 1000000000 {
		bandwidthLeftStr = formatSizePixeldrain(float64(bandwidthUsedPaid), 3)
		errMsg := fmt.Sprintf("Limit is exceeded because less than 1 GB, please wait 24 hours. Here's your bandwidth left: %s", bandwidthLeftStr)
		return "", "", errors.New(errMsg)
	}
	file_name, ok := response["name"].(string)
	if !ok {
		return "", "", errors.New("failed to cast file_name to string")
	}
	filename = file_name
	return filename, dlLink, err
}
