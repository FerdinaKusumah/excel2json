# Convert Excel Or Csv To Json Easily

## Read excel `xlsx` from remote source

### Supposed you have excel like this
![Excel Image](https://raw.githubusercontent.com/FerdinaKusumah/excel2json/master/image/excel_image.png)

### Need to convert to this
```json
[{
    "Profit": "-213.25",
    "ShippingCost": "35",
    "UnitPrice": "38.94"
}, {
    "Profit": "457.81",
    "ShippingCost": "68.02",
    "UnitPrice": "208.16"
}, {
    "Profit": "46.71",
    "ShippingCost": "2.99",
    "UnitPrice": "8.69"
}]
```

```go
import (
    "encoding/json"
	"fmt"
	"github.com/FerdinaKusumah/excel2json"
	"log"
)

func main() {
	var (
		result    []*map[string]interface{}
		err       error
		url       = "https://www.wisdomaxis.com/technology/software/data/for-reports/Data%20Refresh%20Sample%20Data.xlsx"
        // select sheet name
		sheetName = "Sheet1"
        // select only selected field
        // if you want to show all headers just passing nil or empty list
		headers   = []string{"Profit", "Shipping Cost", "Unit Price"}
	)
	if result, err = excel2json.GetExcelFileUrl(url, sheetName, headers); err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}
	for _, val := range result {
		result, _ := json.Marshal(val)
		fmt.Println(string(result))
	}
}
```

## Read excel `xlsx` from local source
```go
import (
    "encoding/json"
	"fmt"
	"github.com/FerdinaKusumah/excel2json"
	"log"
)

func main() {
	var (
		result    []*map[string]interface{}
		err       error
		path      = "./Data Refresh Sample Data.xlsx"
		// select sheet name
		sheetName = "Sheet1"
        // select only selected field
        // if you want to show all headers just passing nil or empty list
		headers   = []string{"Profit", "Shipping Cost", "Unit Price"}
	)
	if result, err = excel2json.GetExcelFilePath(path, sheetName, headers); err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}
	for _, val := range result {
		result, _ := json.Marshal(val)
		fmt.Println(string(result))
	}
}
```

### Supposed you have csv data like this
![Csv Image](https://raw.githubusercontent.com/FerdinaKusumah/excel2json/master/image/csv_image.png)

### Need to convert to this
```json
[{
    "humidity": "17.9279166666667",
    "sound": "1072.35"
}, {
    "humidity": "39.8218442118227",
    "sound": "1683.42557471264"
}, {
    "humidity": "50.5238063660478",
    "sound": "1542.53413589063"
}, {
    "humidity": "28.7329707285052",
    "sound": "1799.60272310065"
}, {
    "humidity": "48.3040024630542",
    "sound": "1452.3183908046"
}]
```

## Read csv from remote source
```go
import (
    "encoding/json"
	"fmt"
	"github.com/FerdinaKusumah/excel2json"
	"log"
)

func main() {
	var (
		result    []*map[string]interface{}
		err       error
		url       = "https://raw.githubusercontent.com/curran/data/gh-pages/senseYourCity/all.csv"
		// select only selected field
        // if you want to show all headers just passing nil or empty list
        headers   = []string{"humidity", "sound"}
		delimited = ","
	)
	if result, err = excel2json.GetCsvFileUrl(url, delimited, headers); err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}
	for _, val := range result {
		result, _ := json.Marshal(val)
		fmt.Println(string(result))
	}
}
```

## Read csv from local source
```go
import (
    "encoding/json"
	"fmt"
	"github.com/FerdinaKusumah/excel2json"
	"log"
)

func main() {
	var (
		result    []*map[string]interface{}
		err       error
		path      = "./all.csv"
		// select only selected field
        // if you want to show all headers just passing nil or empty list
        headers   = []string{"humidity", "sound"}
		delimited = ","
	)
	if result, err = excel2json.GetCsvFilePath(path, delimited, headers); err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}
	for _, val := range result {
		result, _ := json.Marshal(val)
		fmt.Println(string(result))
	}
}
```

### Contributors are welcome !!
