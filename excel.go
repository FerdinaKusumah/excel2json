package excel2json

import (
	"bytes"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strings"
)

// parseExcelFileData to read all data excel
// data is byte data from http client
func parseExcelFileData(data []byte, sheetName string) ([]map[string]interface{}, error) {
	var (
		headers []string
		result  []map[string]interface{}
		wb      = new(excelize.File)
		err     error
	)
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
	return result, nil
}

