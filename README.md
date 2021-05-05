# gomix

サイトURL:http://gomix.work

## 概要

書籍「みんなのGo言語」をベースに様々なツールを作成したアプリ gomix  
数値変換・メモJSONファイル作成・コマンドラインツールなど、  
様々なツールを作成したアプリをGCEへデプロイ  

GCEの構成は下記URLのインフラ構成図を御覧ください。  
https://drive.google.com/file/d/1JReE-3uQj2W4v8GXK0GqYEqHXG_EUd2e/view?usp=sharing

GAE+CloudSQLでもデプロイしました。(Datastore導入予定)  
サイト:https://hip-cyclist-310707.df.r.appspot.com  
ソースコード:https://github.com/keigooba/gomixGAE.git

## 初期設定

<p>Go ModuleやMakefileなど、詳細はSETTING.mdに記載</p>

## 機能

1. 2進数・16進数変換  
<strong>入力した10進数+19の20個までの結果が表示されます</strong>
   - 正規表現・全角 → 半角変換・文字列 → 数値変換、contextによるタイムアウト処理
   - goroutineを用いた5つ並列処理・mutexによる処理のロック
   - エラーメッセージの画面表示・テストケースの作成
   - syncによる同期処理、処理時間計測
2. メモ・JSONファイル作成  
<strong>入力した内容がメモとして記録されます</strong><br>
<strong>JSON形式で入力したファイルはDBに保存できます</strong>
   - osパッケージのみを用いたメモ帳の作成
   - ディレクトリの自動作成・Unix タイムを用いてファイル名を動的変更
   - ioutil.ReadDir のディレクトリ検索・path.Match のファイル名一致
   - JSON形式で入力したファイルをORM(gorm)を使用してDB保存
   - バッファリング、go-isattyによる出力先の判別、ハッシュ値の生成
   - humanizeを用いたバイト数のログ出力
3. コマンドラインツールによるORMの操作  
<strong>DBテーブルをコマンドラインによって操作できます</strong><br>
<strong>コマンドラインの操作は[バイナリファイル] [-h]or[-help]でご確認下さい</strong><br>
   - メモテーブルの作成・検索・編集・削除
4. ファイル出力  
<strong>メモに保存されたtxtファイルが画面出力されます</strong><br>
   - ゴルーチンとreflectを用いた画面出力

## 技術

1. Go1.16.2
2. Bootstrap4.5.0
3. jQuery3.5.1
4. MySQL8.0.23

## 学習記録・作成物

1. [みんなのGo言語](https://www.amazon.co.jp/%E6%94%B9%E8%A8%822%E7%89%88-%E3%81%BF%E3%82%93%E3%81%AA%E3%81%AEGo%E8%A8%80%E8%AA%9E-%E6%9D%BE%E6%9C%A8-%E9%9B%85%E5%B9%B8/dp/4297107279) 3月26日~4月24日  
2. [現役エンジニアが教える、手を動かして学ぶGoogle Cloud Platform(GCP) 入門](https://www.udemy.com/share/1024VCAEYcdVpUQHQD/) 4月25日~5月4日
