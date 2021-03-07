# 組み合わせ

Permutations の列挙

## 私の解答

- 解法: 素直に DFS

早く再帰からループで書けるようになれという気持ち。
今回もノードから戻る際に (図中"ここ" )、スライスで意図しない挙動を起こした(`[3, 3]`になった)のでギブアップ。

```dot
digraph G {
  A [
    label="[ ]\n[1,2,3]"
  ]
  B [
    label="[1]\n[2, 3]"
  ]
  C [
    label="[2]\n[1, 3]"
  ]
  D [
    label="[3]\n[1, 2]"
  ]
  BB [
    label="[1, 2]\n[3]"
  ]
  BC [
    label="[1, 3]\n[2]"
  ]
  BBB [
    label="[1, 2, 3]\n[ ]"
  ]
  BCB [
    label="[1, 3, 2]\n[ ]"
  ]
  A -> {B, C, D}
  B -> {BB, BC}
  BB -> BBB
  BC -> BCB
  BBB -> B [ label="ここ" ]
}
```

## 解答例１

- とりあえずコピーする

`append(left[:i], left[i+1:]...)` のようなコードは必ずバグを発生させている気がするので、分割して 1 つ 1 つ足していくのが無難そう。