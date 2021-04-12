# gomix

## 概要

書籍「みんなのGo言語」をベースに様々なツールを作成したアプリ gomix  
数値変換、メモ帳作成など、様々なツールを作成したアプリを、GCE へデプロイ(予定)

動画で機能をまとめて解説しています。(予定)

## 初期設定

<p>Go ModuleやMakefileなど、詳細はSETTING.mdに記載</p>

## 機能

1. 2 進数・16 進数変換
   - 正規表現・全角 → 半角変換・文字列 → 数値変換
   - goroutine を用いた 5 つ並列処理・mutex による処理のロック
   - エラーメッセージの画面表示・テストケースの作成
   <ul>
    <li><strong>入力した10進数+19の20個まで2進数・16進数変換した結果が表示されます</strong></li>
   </ul>
2. メモ・JSONファイル作成
   - osパッケージのみを用いたメモ帳の作成
   - ディレクトリの自動作成・Unix タイムを用いてファイル名を動的変更
   - ioutil.ReadDir のディレクトリ検索・path.Match のファイル名一致
   - JSON形式で入力したファイルをgormを使用し、DB保存
   - バッファリング、go-isattyによる出力先の判別
   <ul>
    <li><strong>入力した内容がメモとして記録されます</strong></li>
    <li><strong>JSON形式で入力したファイルはDBに保存できます</strong></li>
   </ul>

## 技術

1. Go1.16.2
2. Bootstrap4.5.0
3. jQuery3.5.1

## 学習記録・作成物

1. [みんなのGo言語](https://www.amazon.co.jp/%E6%94%B9%E8%A8%822%E7%89%88-%E3%81%BF%E3%82%93%E3%81%AA%E3%81%AEGo%E8%A8%80%E8%AA%9E-%E6%9D%BE%E6%9C%A8-%E9%9B%85%E5%B9%B8/dp/4297107279) 3月26日~現在
