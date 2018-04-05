package model

var FakeJson1 = `{
	"fields": [{
		"name": "_bkt"
	}, {
		"name": "_cd"
	}, {
		"name": "_eventtype_color"
	}, {
		"name": "_indextime"
	}, {
		"name": "_raw"
	}, {
		"name": "_serial"
	}, {
		"name": "_si"
	}, {
		"name": "_sourcetype"
	}, {
		"name": "_time"
	}, {
		"name": "entity"
	}, {
		"name": "eventtype"
	}, {
		"name": "host"
	}, {
		"name": "index"
	}, {
		"name": "linecount"
	}, {
		"name": "log"
	}, {
		"name": "punct"
	}, {
		"name": "source"
	}, {
		"name": "sourcetype"
	}, {
		"name": "splunk_server"
	}, {
		"name": "splunk_server_group"
	}],
	"highlighted": {},
	"init_offset": 0,
	"messages": [],
	"preview": true,
	"results": [{
		"_bkt": "events-bf2b717b~0~DB09035E-93D7-44EA-839D-67599BF962D6",
		"_cd": "0:4",
		"_indextime": "1522955980",
		"_raw": "{\"entity\":\"test_api\",\"log\":\"This is my first Nova event\"}",
		"_serial": "0",
		"_si": ["si-b24c61c6-69bf959f5d-m8j7k", "events-bf2b717b"],
		"_sourcetype": "httpevent",
		"_time": "2018-04-05T19:19:40.000+00:00",
		"entity": ["test_api", "test_api"],
		"host": "si-b24c61c6:8088",
		"index": "events-bf2b717b",
		"linecount": "1",
		"log": "This is my first Nova event",
		"punct": "{\"\":\"\",\"\":\"_____\"}",
		"source": "curl",
		"sourcetype": "httpevent",
		"splunk_server": "si-b24c61c6-69bf959f5d-m8j7k"
	}]
}`

func GetFakeJson1() (string) {
	return FakeJson1
}