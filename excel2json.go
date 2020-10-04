package excel2json

// GetFileUrl to read all data excel
// url is string uri
// sheetName is sheet name in excel file
func GetFileUrl(url, sheetName string) ([]map[string]interface{}, error) {
	var (
		result   []map[string]interface{}
		err      error
		byteFile []byte
	)
	if byteFile, err = getFileUrl(url); err != nil {
		return nil, err
	}
	if result, err = parseExcelFileData(byteFile, sheetName); err != nil {
		return nil, err
	}
	return result, nil
}
