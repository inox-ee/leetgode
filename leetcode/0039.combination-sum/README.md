# 和がある数となる組み合わせ列挙

配列 `nums` から、重複を許してその和がある数 `target` となるような組み合わせを見つける

## 私の解答

_0022.generate-parentheses_ と同様に再帰で解いた。

### WA となる原因の調査

一部状況で、L15 `*ans = append(*ans, append(tmp, v))` の際に、既存の `ans` が書き換えられる現象が発生。
これは `tmp` が繰り返し利用されることで、容量が肥大化し、意図せぬポインタのバグが生じていることが原因。

## 解答例１

上記の問題を解決するため、更新の際には同時に capacity も更新する。
具体的には、容量を長さと同じにしておくことで、append が行われた時にコピーを作らせることができる。

```Go
solution = solution[:len(solution):len(solution)]
```

Go では、`a[low : high]` に加え、`a[low : high : max]` と記述することで、`max - low` に容量を指定できる。

### 参考

- 公式: <https://golang.org/ref/spec#Slice_expressions>
- stackoverflow: <https://stackoverflow.com/questions/27938177/golang-slice-slicing-a-slice-with-sliceabc> / <https://stackoverflow.com/questions/12768744/re-slicing-slices-in-golang>
