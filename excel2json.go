package main

import "fmt"

// GetExcelFileUrl to read all data excel
// url is string url file name
// sheetName is sheet name in excel file
// Headers is array string that is want to select specific data
func GetExcelFileUrl(url, sheetName string, headers []string) ([]map[string]interface{}, error) {
	var (
		result   []map[string]interface{}
		err      error
		byteFile []byte
		keyName  = fmt.Sprintf(`%s||%s`, url, sheetName)
	)
	if byteFile, err = getFileUrl(url); err != nil {
		return nil, err
	}
	if result, err = parseExcelFileData(byteFile, keyName, sheetName); err != nil {
		return nil, err
	}
	if len(headers) > 0 {
		result = filterDataHeaders(mapHeaders(headers), result)
	}
	return result, nil
}
