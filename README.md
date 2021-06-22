# `Set-Cookie`ヘッダーの`Secure`属性の動作検証

以下のコマンドでサーバーをビルドする。

```console
go build
```

##  動作検証の手順

1. `./server`を実行してサーバーを起動する
2. ブラウザのシークレットモードで`http://localhost:4000/home`にアクセスする
3. Loginリンクをクリックする
4. GitHubのトップページにリダイレクトされるのでブラウザの「戻る」ボタンで戻る
5. ブラウザのリロードボタンを押してページを更新する
6. Cookieが付与されているとAfter loginと表示される

## 結果

- Safari ... Cookieが付与されない（Before login表示のまま）
- Google Chrome ... Cookieが付与される（After loginと表示される）

Google Chromeはlocalhostに対して`Set-Cookie`ヘッダーの`Secure`属性が無視される?

動作検証に使用したバージョンはそれぞれ以下の通り
- Safari Version 14.1 (16611.1.21.161.6)
- Google Chrome Version 91.0.4472.114 (Official Build) (x86_64)
