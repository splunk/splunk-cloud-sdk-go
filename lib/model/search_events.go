package model

// SearchEvents is a type representing a search job
type SearchEvents struct {
	Preview     bool          `json:"preview"`
	InitOffset  int           `json:"init_offset"`
	Messages    []interface{} `json:"messages"`
	Results     []*Result
	Fields      []map[string]interface{} `json:"fields"`
	Highlighted map[string]interface{}   `json:"highlighted"`
}

// Result contains information about the search
type Result struct {
	Bkt          string   `json:"_bkt"`
	Cd           string   `json:"_cd"`
	IndexTime    string   `json:"_indextime"`
	Raw          string   `json:"_raw"`
	Serial       string   `json:"_serial"`
	Si           []string `json:"_si"`
	SourceType1  string   `json:"_sourcetype"`
	Time         string   `json:"_time"`
	Entity       []string `json:"entity"`
	Host         string   `json:"host"`
	Index        string   `json:"index"`
	LineCount    string   `json:"linecount"`
	Log          string   `json:"log"`
	Punct        string   `json:"punct"`
	Source       string   `json:"source"`
	SourceType2  string   `json:"sourcetype"`
	SplunkServer string   `json:"splunk_server"`
}
