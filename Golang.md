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

## 配列のポインタ変数から各 index にアクセスする

```Go
var arr = []int{1,2,3,4}
var parr = &arr

fmt.Println((*parr)[1])  //  2
```

## Slice に関する操作

<https://i-beam.org/2019/12/09/go-slice-tricks/> がメモリ確保の図付きで書かれているので参照されたい。

### append に関するよく分からない(なかった)挙動

append を使っていて、「`append(a, b...)` は `a` を変更しうるのか？」という疑問が生じた。
これに対する解は、「変更することもあればそうでない時もある」となる。
`append()` はコピー先 `a` にコピー元 `b` の全ての要素をコピーする。この時、連結後の長さが `cap(a)` を超える場合は、**新たな slice が作成されて `a` `b` ともにコピーされる。**
通常 `arr = append(arr, 2)` などと書く場合が多いが、**コピー先の slice をスライスした場合には確実に破壊的変更が発生する** ので気を付けなければならない。

```Go
arr := []int{1, 2, 3, 4, 5}
arr2 := append(arr[:2], 10)
/*
この時

arr2 = [1, 2, 10]
arr  = [1, 2, 10, 4, 5]

となる。
*/
```

### 比較

固定長配列の比較は `a == b` で可能だが、slice はそうではない。
`import "reflect"` から `reflect.DeepEqual(a, b)` とする。(-> [docs](https://golang.org/pkg/reflect/#DeepEqual))
なお順序も比較対象なので注意。

## スライスのソート

<https://pepese.github.io/blog/golang-sort/> がよくまとまっていた。

基本的に `sort.Sort()` は、第一引数に `interface` を、第二引数に `less function`[^less-func] を与える。
なお Slice や IntSlice など、代表的なソート対象はもとから用意されているものも多いので [docs](https://golang.org/pkg/sort/) を調べてみることをすすめる。

[^less-func]: `less` って何？
