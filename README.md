# TweetName
天気予報を取得してTwitterの名前欄に反映するアプリケーション

### 内容物
気象庁の天気予報を取得して、適当に整形してツイッターの名前を生成するpython2.7スクリプト  
＋  
実行時引数を絵文字変換してからツイッターの名前を変えるGo言語  

### 利用データ
出典：[気象庁ホームページ](http://www.jma.go.jp/jma/)  
気象庁ホームページの天気予報は出典を明記することで[二次利用可能](http://www.jma.go.jp/jma/kishou/info/coment.html)です。  
Twitterの名前欄に出典を明記できないので、こちらで出典を記載する形をとらせていただきます。

### 使い方
1. Twitter Developersに登録する(説明省略)
1. Go言語の実行環境を整える
    + Goのコンパイルをするために下記をインストールする  
      go get github.com/mrjones/oauth  
      go get golang.org/x/text/encoding  
      go get gopkg.in/kyokomi/emoji.v1  
1. Githubからソースコードを任意のフォルダに展開する
1. 上記フォルダでGoのソースをコンパイルする
1. oauth.jsonを作成する
    + oauth_sample.jsonをリネームして、登録済みの内容に書き換える
1. python 2.7の実行環境を整える
    + そもそも2.7なんて使わない3.xな人はソースコードの.encode('utf-8')を消せば行けるはず。たぶん
1. weather.pyのkenCodeを気象庁のページに合わせて変える
    + デフォルトでは[群馬県の天気予報](http://www.jma.go.jp/jp/yoho/315.html)を表示するため、URLに合わせてkenCodeを`315`としている
      例えば東京の方は`319`に変更すること
1. weather.pyのnameFmtを自分用に変える
    + `-tenki-`を天気予報の記述に置換する仕組みである
    + 天気予報の記述が不満な人はgetName関数を好きなように書き換える
      https://www.webpagefx.com/tools/emoji-cheat-sheet/
1. ネットにつながったPCで下記の処理を実行する
    $ cd 解凍済みフォルダ
    $ python weather.py | TweetName
    ※Windowsの場合は`TweetName.exe`に変更
1. Twitterで名前を確認する
    + あとはcronやタスクスケジューラ仕込むなどご自由に。

### 開発経緯
1. [Twitterの名前を5分毎に東京の天気☼☂☃と連動させるサーバレスプログラムを書いたら色々知らないことが出てきた話](https://qiita.com/issei_y/items/ab641746be2704db98be)を読む
2. 週末の天気が一目でわかると便利そうだと思いつく
2. フリーのAPIだと5日しか取れなかったり精度が悪かったりという結果を見たので、気象庁のデータを使用する
3. AWSでサーバレスするお金を掛けたくないので、前から持ってたConoHaサーバで適当なコードを書く
4. サーバ環境のせいで、慣れないpython2.7の文字コードに苦戦する
5. サーバにtweepy入れようとして失敗したので、実績のある[Go言語](https://github.com/payaneco/GutenJapAlice)でツイートするよう妥協する
6. cronで5:15, 11:15, 17:15に処理を仕込んで出来上がり

### 雑感
このreadmeを書いてるだけで、普段プログラミングしない人が利用するにはハードルが高い(相当頑張らないと多分無理)と思いました。  
  
自分用プログラムなので、気象庁の天気予報で先頭に表示されている群馬県南部の天気を取得しています。  
群馬県北部や小笠原諸島の天気が欲しい方は、自力でpythonスクリプトを書き換えてください。  
  
作った当初は誰でも利用できる連携サービスにしようかと思いましたが、あまりデバッグしてないサービスで他者のアカウント書込み権限を預かるなんて怖すぎてできません。  
  
もしほかに名前変更サービス作りたい方がいれば、天気予報を追加するだけでなく、下記の例のようなリスト形式で日付と名前を変更できる名前スケジューラにすると価値が上がるんじゃないかと思いました。  
  
    2018/0X/XX 〇〇＠ブルベ300-tenki- 
    2018/0Y/YY 〇〇＠ブルベ400-tenki- 
    2018/01/01 ××　禁酒-count-日目
  
こんな感じで、特定日付に名前を変える仕様や、自動カウントアップする仕様のサービスは面白いかなと。  
(例えば`2018/01/01 ××　禁酒-count-日目`の人は、2018/01/03には名前が`××　禁酒３日目`となるイメージで書いています)  
  
私自身は名前欄にスケジュールを表記しないので、作る気もない発案ですが、作りたい方がいれば挑戦してみてください。

### 免責
MITライセンスなので、良識の範囲内でご利用ください。  
ただし気象庁の天気予報をスクレイピングして加工する都合上、天気予報のレイアウトやHTMLの内部構造が変わった時にエラーとなります。  
また、過度に気象庁のホームページへアクセスすることもご遠慮願います。  
  
利用による不具合や被害は一切保証しません。  
機能追加の要望も受け付けません。
