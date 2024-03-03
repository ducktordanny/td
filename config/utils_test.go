package config

import (
	"testing"
	"time"

	"github.com/ducktordanny/td/flags"
)

// TODO: Check testing best practices

var time1, _ = time.Parse(time.RFC3339, "2024-02-22T20:38:48.000000+01:00")
var time2, _ = time.Parse(time.RFC3339, "2024-01-22T20:38:11.000000+01:00")
var time3, _ = time.Parse(time.RFC3339, "2024-02-21T20:30:41.000000+01:00")
var timePointer = time.Now()

var listMock = []TodoModel{
	{
		Id:        "abc1234",
		Content:   "Test1",
		Resolved:  false,
		CreatedAt: time1,
	},
	{
		Id:        "abc1235",
		Content:   "Test2",
		Resolved:  true,
		CreatedAt: time2,
	},
	{
		Id:         "abc1236",
		Content:    "Test3",
		Resolved:   false,
		CreatedAt:  time3,
		ResolvedAt: &timePointer,
	},
}

func TestGenerateUniqueSha(t *testing.T) {
	resultA := GenerateUniqueSha()
	resultB := GenerateUniqueSha()
	if resultA == resultB {
		t.Errorf("Two results of `GenerateUniqueSha()` should not have the same value. A: %s, B: %s", resultA, resultB)
	}
	if len(resultA) != 7 {
		t.Errorf("Result's length should be 7, got: %d (for '%s')", len(resultA), resultA)
	}
}

func TestGetListBasedOnScopeLocal(t *testing.T) {
	_GetConfigPath := GetConfigPath
	defer func() {
		GetConfigPath = _GetConfigPath
	}()
	GetConfigPath = func() string {
		return "../mocks/config.json"
	}

	_GetLocalPath := GetLocalPath
	defer func() {
		GetLocalPath = _GetLocalPath
	}()
	GetLocalPath = func() string {
		return "/Users/test/b"
	}

	conf := ReadConfig()
	result := GetListBasedOnScope(flags.LocalScope, &conf)
	if len(*result) != 3 {
		t.Error("Result should have length of 3.")
	}
	if (*result)[0].Content != "I should implement this and that" {
		t.Error("First item should have content of 'I should implement this and that'")
	}

}

func TestGetListBasedOnScopeGlobal(t *testing.T) {
	_GetConfigPath := GetConfigPath
	defer func() {
		GetConfigPath = _GetConfigPath
	}()
	GetConfigPath = func() string {
		return "../mocks/config.json"
	}

	conf := ReadConfig()
	result := GetListBasedOnScope(flags.GlobalScope, &conf)
	if len(*result) != 2 {
		t.Error("Result should have length of 2.")
	}
	if (*result)[0].Content != "Test" {
		t.Error("First item should have content of 'Test'")
	}
}

func TestGetIndexOfListBySha(t *testing.T) {
	index := GetIndexOfListBySha(listMock, "abc1234")
	if index != 0 {
		t.Error("ID 'abc1234' should return 0")
	}
	index = GetIndexOfListBySha(listMock, "abc1235")
	if index != 1 {
		t.Error("ID 'abc1235' should return 1")
	}
	index = GetIndexOfListBySha(listMock, "abc1236")
	if index != 2 {
		t.Error("ID 'abc1236' should return 2")
	}
}

func TestGetLocalModelByPath(t *testing.T) {
	_GetConfigPath := GetConfigPath
	defer func() {
		GetConfigPath = _GetConfigPath
	}()
	GetConfigPath = func() string {
		return "../mocks/config.json"
	}

	conf := ReadConfig()
	result := GetLocalModelByPath(&conf, "/Users/test/a")
	if (*result).Path != "/Users/test/a" {
		t.Error("Path should be '/Users/test/a'")
	}
}
