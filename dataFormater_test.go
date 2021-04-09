package risk

import (
	"testing"
)

// 应用配置
var appConfig string = `
{
    "column_config": [
        {
            "key_name": "deduct_type",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "ip3",
            "func_name": "subIp",
            "args": [
                "ip",
                0,
                "3"
            ]
        },
        {
            "key_name": "referer",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "os",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "pack_name",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "website",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "token",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "ext_id",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "androidid",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "oaid",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "idfa",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "imei",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "mac",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "create_time_ms",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "create_date",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "province_id",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "city_id",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "user_agent",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "change_point",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "ip_type",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "ip",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "ad_source",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "ad_type",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "content_type",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "product_price",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "item3",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "item2",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "item1",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "content_id",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "hos_city",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "hos_province",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "hospital_id",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "gps",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "device_type",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "device_id",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "msg_id",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "order_money",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "order_id",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "channel_name",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "sys",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "register_date",
            "func_name": "",
			"type" :"string",
            "args": null
        },
        {
            "key_name": "uid",
            "func_name": "",
			"type" :"string",
            "args": null
        }
    ],
    "strategy_id_list": [
        "123",
        "124",
        "125",
        "126",
        "127",
        "128",
        "129",
        "130",
        "131",
        "132",
        "133",
        "134",
        "135",
        "137",
        "138",
        "139",
        "140",
        "141",
        "142",
        "143",
        "144",
        "145",
        "146",
        "147",
        "148",
        "149",
        "150",
        "151",
        "152",
        "153",
        "154",
        "155",
        "156",
        "157",
        "158",
        "159",
        "160",
        "161",
        "162",
        "165"
    ],
    "rule_config": [
        {
            "rule_id": "16",
            "app_id": "4"
        },
        {
            "rule_id": "15",
            "app_id": "4"
        },
        {
            "rule_id": "21",
            "app_id": "4"
        }
    ],
    "version": "v2.1"
}
`

// 策略配置
var StrategyConfig string = `
{
    "id": 16,
    "app_id": 4,
    "weight": 4,
    "trigger": 8,
    "status": 1,
    "rule_filter": {
        "theme": "any",
        "rule": [
            {
                "rule_id": 57
            },
            {
                "rule_id": 56
            }
        ]
    }
}
`

// 字段配置
var DataFormatterConfig string = `
[
    {
        "key_name": "ip3",
        "func_name": "subIp",
        "args": [
            "ip",
            0,
            "3"
        ],
        "type": "string"
    },
    {
        "key_name": "ip",
        "func_name": "",
        "args": null,
        "type": "string"
    },
    {
        "key_name": "hospital_id",
        "func_name": "",
        "type": "int",
        "args": null
    },
    {
        "key_name": "create_time_ms",
        "func_name": "",
        "type": "int",
        "args": null
    },
    {
        "key_name": "create_date",
        "func_name": "",
        "type": "int",
        "args": null
    },
    {
        "key_name": "uid",
        "func_name": "",
        "type": "int",
        "args": null
    }
]
`

// 简单指标原子配置
var SimpleQuotaConfig string = `
{
    "strategy_atom_id": 163,
    "condition": "count",
    "search_key": [
        "hospital_id"
    ],
    "count_key": "uid",
    "system_sign": 0,
    "version": "v2.1",
    "app_id": 104,
    "time_shard": 86400000,
    "filter": {
        "theme": "any",
        "rule": [
            {
                "source": "uid",
                "condition": "notEmpty",
                "target": 11111
            }
        ]
    }
}
`

// 复合指标配置
var MultiQuotaConfig string = `
{
    "complex_quota_id": 35,
    "system_sign": 0,
    "version": "v2.1",
    "app_id": 107,
    "filter": [],
    "formula": "#189\/#181",
    "quota_map": {
        "#181": {
            "app_id": 107,
            "strategy_id": 181
        },
        "#189": {
            "app_id": 107,
            "strategy_id": 189
        }
    }
}
`

// 规则配置
var ruleConfig string = `
{
    "rule_id": 94,
    "warning_point": 0.2,
    "forbidden_point": 0.5,
    "filter": {
        "theme": "",
        "rule": [
            {
                "source": "",
                "condition": "",
                "target": ""
            }
        ]
    },
    "strategy_config": {
        "id": 35,
        "type": "multi"
    }
}
`

var RawData string = `
{
	"uid": 14625461,
	"ad_type": 21,
	"sys": 1,
	"device_id": 203310581,
	"hospital_id": "108720",
	"content_id": 597322,
	"content_type": 1,
	"ip": "61.50.119.46",
	"change_point": 5,
	"create_date": "2021-04-09 15:06:20",
	"province_id": 1,
	"city_id": 1,
	"website": "app",
	"ip2city": "1",
	"ip2province": "1",
	"ext_id": "174127",
	"deduct_type": 1
}
`

func TestDemo(t *testing.T) {
	Text()
}

func TestConfig(t *testing.T) {
	t.Log("start")
	rd := RiskData{}
	err := rd.LoadConfig(DataFormatterConfig)
	if err != nil {
		t.Error(err)
	}
	t.Log(rd)
	t.Log("end")
}

func TestData(t *testing.T) {
	t.Log("start")
	rd := RiskData{}
	err := rd.LoadConfig(DataFormatterConfig)
	if err != nil {
		t.Error(err)
	}

	err = rd.LoadData(RawData)
	if err != nil {
		t.Error(err)
	}

	t.Log(rd)
	t.Log("end")
}

func TestJsonDecode(t *testing.T) {
	t.Log("start")
	data := `
{
    "uid": 14625461,
    "ad_type": 21,
    "sys": 1,
    "hospital_id": "108720"
}
`
	tpl := map[string]interface{}{}
	JsonDecode(data, &tpl, true)
	t.Logf("%#+v \n", tpl)
}
