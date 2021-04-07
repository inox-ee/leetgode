# 最大部分配列問題

## 私の解答

- 解法：全探索

$O(n^2)$ なので当然 `Time Limit Exceeded N/A N/A`

動的計画法のコツが分からん

## 解答例１

- 解法：動的計画法

最大/最小問題では、動的計画法を検討してみる。
本問題では、評価値に負の数が含まれていることがネックだが、この「最大部分配列問題」は _Kadane's algorithm_ が有効である。

### Kadane's algorithm

貪欲法の一種。

簡単のため、次の配列を考える。
$nums := [-2, 1, -3, 4, ...]$
(正確には違うが) `current_subarray` が $[-2, 1, -3]$ とすると、`next_subarray` は **`current_subarray + [4]` と `[4]` の比較**により求められる。前者 < 後者 の場合、その部分列の後ろにいかなる値を加えても最大部分列とはならないことがわかる。

### ポイント？

$a[i]$ と $a[i+1]$ の大小比較を行うのではなく、$a[i+1]$ の条件を考えるべき。
初め動的計画法を思いついた際は、$a[i]$ と要素 $i+1$ の和と $a[i]$ の大小関係を考えてしまったので求められなかった。あくまで $a[i+1]$ になりうる選択肢を考えよう。

## 結果

    Runtime: 128 ms, faster than 6.55% of Go online submissions for Maximum Subarray.
    Memory Usage: 4.3 MB, less than 5.08% of Go online submissions for Maximum Subarray.

なんでぇ…
