# Real World HTTP
読んでみた自分用のまとめ（勘違いなどあるかも）

## 第1章

- dumpはンピュータの記憶装置（メモリやストレージなど）に記録された内容をまとめて表示すること
- 301はリクエストされたページが別の場所に移動した時に使う.
- 302はモバイル専用サイトにジャンプしたり, 麺店アンスページを表示したりする一時的な移動に使う.
- 303はログインページを使ってログインした後に, 元のペー位に飛ぶ場合などに使う.
- クライアントは300番台のステータスを受け取って, リダイレクトする時は, レスポンスヘッダーのLocationヘッダーの値を見て再度リクエストを送信し直す.
- headerのx-tokenなどのxには独自headerという意味がある

## 第2章
- クッキーはサーバー側でレスポンスヘッダーにクッキーで保存する情報を格納してフロントで受け取って, ブラウザに保存する. フロントがリクエストするときにクッキーの内容をヘッダーに格納してサーバー側が読み取る
- クッキーに保存する内容は「無くなっても問題のない情報」のみにして「システムで必要となるID」や「書き換えられると誤動作につながるようなセンシティブな情報」はなるべく入れないようにする
- オリジン：スキーム・ドメイン・ポートのことで, これが異なるものは別のオリジンとして扱われる. 別のオリジンと通信するときはCORSが必要になる.
- BASIC認証はユーザ名とパスワードをbase64エンコーディングして認証する方式. SSL/TLS通信を使っていない状態だとユーザ名とパスワードが漏洩してしまう.
- 色々なキャッシュの方法（詳しくは本で）

## 第4,5章
- リクエストヘッダーのConnection: Keep-Aliveによって, 連続したリクエストの時に接続を再利用するようになるため, HTTPの下のレイヤーであるTCP/IPの通信が効率化して通信が高速化する
- ハッシュ値は暗号ではないが, 非常に使いやすい性質をいくつか持っており, TLSでもよく用いられることがある
- 1バイトでもデータに差異があればハッシュ値が変わるので, 内容の同一性を判断する用途でよく使われる
- 共通鍵方式は暗号化と復号化に同じ鍵を使う
- 公開鍵方式は公開鍵を公開して秘密鍵は他の人に知られないように管理する. 公開鍵で暗号化して, 秘密鍵でのみ復号できるというようなイメージ
- デジタル署名は, 例えば秘密鍵を持つサーバーが公開鍵を持つユーザにファイルを送信したいとする. そのまま送信して改竄されても改竄されたかどうかが不明なので, 送信したいファイルをハッシュ化して秘密鍵で暗号化したデジタル署名をファイルに添付する. 受け取ったユーザは公開鍵で復号してファイルの中身が同じかを判断して改竄されたかを判断できる
- 鍵交換（共通鍵を傍受されないための方法）のアルゴリズムの1つであるDH鍵交換アルゴリズムは, 鍵そのものを交換せずにクライアントとサーバーでそれぞれ鍵の材料を作り, 互いに交換し合って, それぞれの場所で計算して同じ鍵を得る. サーバーはサーバー独自の材料＋クライアントでも共通の材料から鍵の材料作って, それ＋共通の材料をクライアントに送り, クライアントはクライアント独自の材料＋共通の材料から鍵の材料を作って, サーバーに送る. サーバーはサーバー独自の材料＋共通の材料＋クライアントから送られてきた鍵の材料から鍵を作成し, クライアントはクライアント独自の材料＋共通の材料＋サーバーから送られてきた鍵の材料から鍵を作成する. ここで生成された鍵2つは同じものになる. 例え通信が傍受されていたとしても, 傍受した側が持っているのは, 共通の材料＋クライアントから送られてきた鍵の材料＋サーバーから送られてきた鍵の材料なので共通鍵が作れない.
- TLSの手順は, まずブラウザがサーバーからサーバーの持つSSLサーバー証明書と公開鍵を取得する. 証明書にあるデジタル署名によってサーバーが改竄されていないことを確認する. ブラウザとサーバーは公開鍵方式による鍵交換によって, 共通鍵を生成する. サーバーとブラウザが共通鍵を取得したら, 共通鍵方式で通信する.

## 第6章
- response.Body.close()でクローズして読み込み終了を明示しなければ, 次のジョブをいつ使い始めれば良いか分からずKeep-Aliveができない.
- 証明書の作成手順は, まずOpenSSLコマンドを使って秘密鍵を作成し, 作成した秘密鍵で証明書署名要求（CSR）を作成する. 署名は署名局で行ってもらう（有料）

## 第10章
- RPCはURLが1つでリクエストの際にサービス名とメソッド名を渡すことで何をするかを指定する.（POSTのみ）
- RPCではサーバーへのリクエストは全てボディーに入れて送信し, ログを見ても同一URLへのアクセスしか確認できない.
- よく見かけるがURLにバージョンが入っているのはRESTとして好ましくない. バージョンやフォーマットはAcceptヘッダーに入れるべき.
- RESTになりきれているか自信がない時は, 「REST-ish API」と呼ぼう.
- PATCHはあまり使っているのを見かけない. OPTIONSは主にCORSのプリフライトリクエストに使われる.
- APIの1秒間あたりの呼び出し回数は10回ぐらいが一般的. Golangでもライブラリがあるので要確認.

## 第12章
- ライブラリがパスとメソッドをもとにどのハンドラを呼び出すか決定するが, この呼び出し先を切り分けるコンポーネント のことはディスパッチャー, ルーター, マルチプレクサーなどの呼び名で呼ばれる.
- 複数のリクエストにまたがったライフサイクルをセッションといい, それを実現する仕組みがクッキー. 初回アクセス時にサーバーからブラウザにクッキーを渡し, 2回目以降のアクセスの時にブラウザはサーバーに対してクッキーを返す. クッキーの中にユーザーIDなどをまとめて入れているのでその情報でAPIを叩ける. もちろんこの場合はクッキーに署名を付与するので改竄はされない. もう1つのやり方として, ログインした時にセッションキーをクッキーで渡す. サーバーはセッション管理用のDBをNoSQLで立てて, セッションキーで情報にアクセスする. 以上のような2つの方法のように, セッションに関連してアクセスしてきたユーザに属する情報を保存する仕組みはセッションストレージという.
- サーバー レスは「自前で管理するサーバーがない」程度の意味合い

## 第13章
- リクエストのたびに毎回DNSサーバーにアクセスするとパフォーマンスが悪いのでキャッシュサーバーを設けてキャッシュすることが多い.
- DNSサーバーにいくつかのIPアドレス（サーバー）を登録しておき, リクエストのたびに呼び出すIPアドレスを変えることでロードバランスするラウンドロビンというやり方もある.
- リバースプロキシはサーバーとクライアントを仲介する位置にあるプロキシのことを言い, 実際に何を行うかはプロキシごとに大きく異なる. コンテンツをキャッシュするものはCDNと言われ, 大量のリクエストを受け, 負荷が均等になるようはいかのサーバーに処理を振り分ける場合はロードバランサーと呼ばれるし, 様々ある. あくまで「リバースプロキシとして実装をされている」というだけ.
- Nginxはウェブサーバーとしての機能以外にもリバースプロキシとしての機能も備えている.
- CDNは, 通信そのものを高速化&安定化させ, ユーザに近い高機能なプロキシサーバーとして機能を提供する.
- DNSでCDNの提供するドメインを参照し, CDNのドメインはIPアドレスから推定してユーザの利用場所に近い場所のホストのIPアドレスを返す.
- コンテンツはCDNのサーバーにキャッシュさせるので通信そのものも高速化する.
- CDNによってオリジンサーバーへのアクセスの集中も避けることができる. また, プロキシサーバーとしてセキュリティ関連の機能を提供するなどプロキシサーバーとしての役割も担うことができる.
- ロードバランサーは管理・稼働しているインスタンスに順番にタスクを回すラウンドロビン, 最終接続時間が遅い順, ハッシュを使って宛先を決める方法などが一般的.
- 正しいタスク分配をするために, インスタンスのヘルスチェックを行い, 正しく動作していないサーバーを割り振り先の対象から外すなどのことも行う. また負荷が一定以上になったらインスタンスを増やして一台あたりの負荷を下げるなどのことをすることもある.
- ロードバランサーは各種サービスへの中継を行うので, ロードバランサーに対して設定変更してサービスを新しいバージョンへ切り替えるといった作業はよく行う. この際いきなり接続を切り替えると利用中のユーザに迷惑がかかるので, 新しいリクエストは新しいバージョンに流して, 古いバージョンへのリクエストは完了するまで待つという制御を行わせる. このことを接続ドレインという.
- ヘルスチェックは指定したHTTPのパスにリクエストに投げることで行う. 結果に問題があるとロードバランサーはリクエストをそのホストにフォーワードしなくなる.
- ヘルスチェックはLiveness ProveとReadiness Proveの2種類がある. Liveness Proveはリクエストが来た時に200だけを返すエンドポイントを用意し, サービスが起動しているかを確認する. Readiness Proveはmサービスを提供できるかどうかを確認します.
- VPCはクラウドサービスの中でプライベートなネットワーク空間を作ることができ, 接続経路を管理したい場合に利用する.VPCの内部にサーバーインスタンスなどを配置する. サブネットを切ったりすることもできる.
- JWTはブラウザのローカルストレージにおいておくのはセキュリティ上望ましくないとされている. そのため, フロント側のサーバーとブラウザ間ではBFFと呼ばれるランダムな認証トークン を利用し, そのBFFでJWTトークン を発行することが多い.
- 分散トレーシング何も分からん
- サーバー内部の時間情報をブラウザで表示するためのHTTPヘッダーもある. Server-TimingやTrailer

# ネットワークはなぜ繋がるのか
- ブラウザでリクエストを作成するが, 送信する機能はない.
- リクエストを作る際, ドメイン名だけでは送信先がわからないので, ドメイン名でDNSサーバーに IPアドレスを問い合わせる
- IPアドレスの仕組みは保存したQiitaと本を要参照
- ブラウザでリクエストが完成すると, OSに送信を頼む
- プロキシはWebサーバーとクライアントの中継役
- DNSサーバーにIPアドレスの代わりに登録する（ちょっと怪しい）
