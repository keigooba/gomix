package main

func main() {
	// エントリーポイントの設定・サーバー起動
	err := StartMainServer()
	if err != nil {
		panic(err)
	}
}