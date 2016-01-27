package repository551_test

import (
	"github.com/go51/repository551"
	"testing"
)

func TestLoad(t *testing.T) {
	m1 := repository551.Load()
	m2 := repository551.Load()

	if m1 == nil {
		t.Errorf("インスタンスの作成に失敗しました。")
	}
	if m2 == nil {
		t.Errorf("インスタンスの作成に失敗しました。")
	}
}

func BenchmarkLoad(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = repository551.Load()
	}
}
