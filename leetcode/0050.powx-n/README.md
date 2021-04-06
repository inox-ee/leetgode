# 累乗

累乗を実装

## 私の解答１

- ふつーに $O(n)$ な実装

当然 TLE

## 私の解答２

- $O(\log n)$ を目指す

指数が奇数のときは $x^n = (x^2)^\frac{n}{2}$、 $x^n = x\times (x^2)^\frac{n-1}{2}$ とすればよい。ループ数が毎回半減するので 計算量 log が実現できる。

$n < 0$ は、場合分けが面倒だったので最後に逆数を取った。

## 結果

    Runtime: 0 ms, faster than 100.00% of Go online submissions for Pow(x, n).
    Memory Usage: 2 MB, less than 18.21% of Go online submissions for Pow(x, n).

メモリがゴミだったんだけど何故だ？再帰もしてないのに
