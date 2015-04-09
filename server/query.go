package server

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Pair struct {
	Name  string
	Value string
}

type Query struct {
	Key   string
	Pairs [49]*Pair
}

var (
	invalidResponse = errors.New("invalid result")
	names           = map[int]string{
		0:  "未知",
		1:  "名字",
		2:  "代码",
		3:  "当前价格",
		4:  "昨收",
		5:  "今开",
		6:  "成交量（手)",
		7:  "外盘",
		8:  "内盘",
		9:  "买一",
		10: "买一量（手）",
		11: "买二",
		12: "买二量（手）",
		13: "买三",
		14: "买三量（手）",
		15: "买四",
		16: "买四量（手）",
		17: "买五",
		18: "买五量（手）",
		19: "卖一",
		20: "卖一量",
		21: "卖二",
		22: "卖二量",
		23: "卖三",
		24: "卖三量",
		25: "卖四",
		26: "卖四量",
		27: "卖五",
		28: "卖五量",
		29: "最近逐笔成交",
		30: "时间",
		31: "涨跌",
		32: "涨跌%",
		33: "最高",
		34: "最低",
		35: "价格/成交量（手）/成交额",
		36: "成交量（手）",
		37: "成交额（万）",
		38: "换手率",
		39: "市盈率",
		40: "未知",
		41: "最高",
		42: "最低",
		43: "振幅",
		44: "流通市值",
		45: "总市值",
		46: "市净率",
		47: "涨停价",
		48: "跌停价",
	}
)

func query(key string) (*Query, error) {
	res, err := http.Get("http://qt.gtimg.cn/q=" + key)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(transform.NewReader(res.Body,
		simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		return nil, err
	}
	bs := string(b)
	begin := strings.Index(bs, "\"")
	end := strings.LastIndex(bs, "~")
	if begin == -1 || end == -1 {
		return nil, invalidResponse
	}
	result := strings.Trim(bs[begin+1:end], "\r\n ")
	if result == "" {
		return nil, invalidResponse
	}
	pairs := strings.Split(result, "~")
	if len(pairs) != len(names) {
		return nil, invalidResponse
	}
	q := &Query{Key: key}
	for i := range q.Pairs {
		q.Pairs[i] = &Pair{Name: names[i], Value: pairs[i]}
	}
	return q, nil
}
