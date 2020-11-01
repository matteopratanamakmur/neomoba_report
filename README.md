# neomoba_report

## 概要
ネオモバのレポート出力

## 環境構築
### chromedriver のインストール
```
$ brew update
$ brew install selenium-server-standalone
$ brew install chromedriver
```

## 認証に失敗する場合
次のように許可を与えてやれば良い
```
% which chromedriver
/usr/local/bin/chromedriver
% cd /usr/local/bin/
% xattr -d com.apple.quarantine chromedriver
```

## user / pass は環境変数から取得
```
% export NEO_USER=<USERNAME>
% export NEO_PASS=<PASSWORD>
```

## 参考ページ
- https://sites.google.com/a/chromium.org/chromedriver/