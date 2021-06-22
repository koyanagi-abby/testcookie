# Cookieの動作検証

以下のコマンドでサーバーをビルドする。

```console
go build
```

- Step 1 ... `./server`を実行してサーバーを起動する
- Step 2 ... ブラウザを開いて`http://localhost:4000/home`にアクセスする
- Step 3 ... Loginリンクをクリックする
- Step 4 ... GitHubのトップページにリダイレクトされるのでブラウザの「戻る」ボタンで戻る
- Step 5 ... Cookieが付与されているとAfter loginと表示される

SafariではlocalhostでもSet-CookieのSecure属性が適用される。ChromeではlocalhostにたいしてSecure属性が無視される。
