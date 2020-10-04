package excel2json

import (
	"bytes"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/patrickmn/go-cache"
	"strings"
	"time"
)

// Create a cache with a default expiration time of 5 minutes, and which
// purges expired items every 10 minutes
var localCache = cache.New(5*time.Minute, 10*time.Minute)

// parseExcelFileData to read all data excel
// data is byte data from http client
func parseExcelFileData(data []byte, keyName, sheetName string) ([]map[string]interface{}, error) {
	var (
		headers []string
		result  []map[string]interface{}
		wb      = new(excelize.File)
		err     error
	)
	// if already data in cache
	if cacheData, found := localCache.Get(keyName); found {
		return cacheData.([]map[string]interface{}), nil
	}
	// open byte data with reader
	if wb, err = excelize.OpenReader(bytes.NewReader(data)); err != nil {
		return nil, err
	}
	// Get all the rows in the Sheet.
	rows := wb.GetRows(sheetName)
	headers = rows[0]
	for _, row := range rows[1:] {
		var tmpMap = make(map[string]interface{})
		for j, v := range row {
			tmpMap[strings.Join(strings.Split(headers[j], " "), "")] = v
		}
		result = append(result, tmpMap)
	}
	// set data to cache
	localCache.Set(keyName, result, 10*time.Minute)
	return result, nil
}
