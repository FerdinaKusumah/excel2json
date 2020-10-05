package excel2json

import (
	"bytes"
	"encoding/csv"
	"strconv"
	"strings"
	"time"
)

func parseCsvFileData(byteFile []byte, delimiter, keyName string) ([]*map[string]interface{}, error) {
	var (
		result    []*map[string]interface{}
		err       error
		headers   []string
		data      [][]string
		delimited int
	)
	// if already data in cache
	if cacheData, found := localCache.Get(keyName); found {
		return cacheData.([]*map[string]interface{}), nil
	}
	// if already delimited
	if delimited, err = strconv.Atoi(delimiter); err != nil {
		delimited = ','
	}
	// initialize reader
	reader := csv.NewReader(bytes.NewReader(byteFile))
	reader.Comma = rune(delimited)
	if data, err = reader.ReadAll(); err != nil {
		return nil, err
	}
	// get headers
	headers = data[0]
	for _, row := range data[1:] {
		var tmpMap = make(map[string]interface{})
		for j, v := range row {
			tmpMap[strings.Join(strings.Split(headers[j], " "), "")] = v
		}
		result = append(result, &tmpMap)
	}
	// set data to cache
	localCache.Set(keyName, result, 10*time.Minute)
	return result, nil
}
