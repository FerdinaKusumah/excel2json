package excel2json

import "strings"

// mapHeaders Convert all headers to map string
// headers is array string
func mapHeaders(headers []string) map[string]string {
	var result = make(map[string]string)
	for _, val := range headers {
		result[strings.Join(strings.Split(val, " "), "")] = strings.Join(strings.Split(val, " "), "")
	}
	return result
}

// filterDataHeaders data based on selected header
// mapHeader is headers that already mapped
// data is data that is already fetch from excel file
func filterDataHeaders(mapHeader map[string]string, data []*map[string]interface{}) []*map[string]interface{} {
	var result []*map[string]interface{}
	for _, val := range data {
		var (
			t = make(map[string]interface{})
			d = *val
		)
		for _, header := range mapHeader {
			if _, ok := d[header]; ok {
				t[header] = d[header]
			}
		}
		result = append(result, &t)
	}
	return result
}
