package main

import "testing"

//单元测试,h函数名需要以Test开头,随后跟大写字母
//参数需要为*testing.T类型，单元测试主要用于测试
//方法或者函数功能是否正常
func TestUnmarshal(t *testing.T) {
	hero, err := unmarshal("hero.json")
	if err != nil {
		t.Error(err)
	}

	if hero == nil {
		t.Error("返回为空")
	}

	if hero.ID != 1 {
		t.Error("return id should be 1,but here is not")
	}
}

//基准测试：函数名必须要以Benchmark开头，参数需要为*testing.B类型
func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("hero.json")
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("hero.json")
	}
}
