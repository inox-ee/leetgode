# Go 知識

## マップ

マップは、第 1 返り値にその値を、第 2 返り値に存在するか否かを返す。
存在しない場合、第 1 返り値は デフォルト値になるので注意。nil ではない。

## 繰り返し

Golang に while は存在しない。
同様の制御文を書きたい場合は、以下の通りにする。

```go
for true {
  ...
}
```

## 再帰

Go で再帰は遅いらしい(末尾でも)。
Go ランタイムのスタックサイズがデフォルトでは小さいのが原因だそう。

ref: <https://ymotongpoo.hatenablog.com/entry/2015/02/23/165341>

## 文字列を１文字ずつ処理

1. `for ... range` を使う場合
   受け取った値は、`int32` 型(要は文字コード)で得られる。4 バイト文字に対応しており、_A_ は 65、_a_ は 97、_あ_ は 12354 など。
   初期値 `res := bytes.Buffer()` に `res.WriteBytes(str[i])` で書き込み、`res.String()` で出力する。
2. `Split()` を使う場合
   受け取った値は `string` 型で得られる。

### [そもそも] Go と文字列

Go では strings の符号化方式として、UTF-8 を使っている。
UTF-8 は、1 byte から 4 byte の可変長データで code point (=文字コード表中のインデックス) を置換している。

ここで問題となるのが、文字列を byte のまま扱うと 4byte 文字にインデックスでアクセスすると正しい文字を得ることができない (JS などの言語で顕著)。

よって Go では*rune*で文字列を扱うのが基本となっている。_rune_ の実体は、 Unicode の code point を表現するための _int32_ である。
以下、引っかかりそうな例。

```go
// len() による文字数カウント
str := "こんにちは、world"
fmt.Println(len(str))        // 23
fmt.Println(len([]rune(s)))  // 11
fmt.Println(utf8.RuneCountInString(s))

// スライス
fmt.Println(str[0:5])                // "こ\xe3\x82"
fmt.Println(string([]rune(s)[0:5]))  // "こんにちは"
```
