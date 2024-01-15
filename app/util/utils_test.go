package util

import (
	"log/slog"
	"reflect"
	"testing"
)

func TestReverseStatus(t *testing.T) {
	testData := []map[string]string{
		{
			"data":     "open",
			"expected": "close",
		},
		{
			"data":     "close",
			"expected": "open",
		},
	}

	for _, v := range testData {
		if exp := ReverseStatus(v["data"]); exp != v["expected"] {
			t.Errorf("Expected %s, got %s", v["expected"], exp)
		}
	}
}

func TestQueryAll(t *testing.T) {
	testData := []map[string]interface{}{
		{
			"data":     map[string]string{"foo": "bar"},
			"expected": []map[string]string{{"foo": "bar"}},
		},
		{
			"data":     map[string]string{"foo": "bar", "bay": "baz"},
			"expected": []map[string]string{{"foo": "bar"}, {"bay": "baz"}},
		},
	}

	for _, v := range testData {
		data, ok := v["data"].(map[string]string)
		if !ok {
			slog.Error("interface revers type error")
		}
		expData, ok := v["expected"].([]map[string]string)
		if !ok {
			slog.Error("interface revers type error")
		}
		if exp := QueryAll(data); !reflect.DeepEqual(exp, expData) {
			t.Errorf("Expected %s, got %s", expData, exp)
		}
	}
}
