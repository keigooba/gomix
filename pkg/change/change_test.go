package change

import "testing"

func TestAdd(t *testing.T) {

	var result int

	// テストケースの検証
	result = Add(1, 2)
	if result != 3 {
		// error表示
		t.Errorf("add failed. expect:%d, actual:%d", 3, result)
	}

	// t.Log, t.Logf でログを出すと `go test -v` と実行したときのみ表示される
	t.Logf("result is %d", result)
}
