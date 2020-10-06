# Convert Excel Or Csv To Json Easily

## Read excel `xlsx` from remote source
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
		sheetName = "Sheet1"
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
		sheetName = "Sheet1"
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