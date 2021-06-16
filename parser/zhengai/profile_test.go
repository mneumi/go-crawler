package zhengai

import (
	"os"
	"testing"

	"github.com/mneumi/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := os.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "何必怀念萌宝")

	if len(result.Items) != 1 {
		t.Errorf("items should contains 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:       "何必怀念萌宝",
		Age:        36,
		Height:     197,
		Weight:     224,
		Income:     "3001-5000元",
		Gender:     "男",
		Xinzuo:     "天秤座",
		Occupation: "测试工程师",
		Marriage:   "未婚",
		House:      "无房",
		Hukou:      "广州市",
		Education:  "博士及以上",
		Car:        "有车",
	}

	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
