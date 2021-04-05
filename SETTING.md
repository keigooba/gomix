# gomix

# 初期設定

1. Go module 導入 go mod init
2. git 導入
3. 各種メトリクス取得 API 導入 golang-stats-api-handler
4. Makefile の作成
5. ログファイルの作成 logrus の導入
6. header,footer の切り分け
7. test ファイルの作成
8. config ファイルの作成
9. golang-lint の導入
10. 自動化 sh ファイルの作成・exec.command で自動実行

# 注意事項

1. runtime.GOOS or //+build を用いて windows でも同様の動作環境で動くようにすること
2. 各種メトリクス取得 API を利用して Munin や Zabbix 等のエージェント経由でメモリや GC の状況をモニタリングすること
