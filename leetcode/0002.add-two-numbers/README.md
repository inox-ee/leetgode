# Add Two Number

## 問題

単連結リストを用いて、整数の和を求める。

## 私の解答

- 解法：単純な再帰
- 計算量：O(max(m,n))

各桁の 2 数＋桁上がりを足してゆく再帰関数を書いた。

## 解答例１

ほぼ同じ。
例外処理を簡単にするため、ダミーヘッドを用意するのが賢い。
