# index 探索

与えられた値 target が整列済み配列 nums にあればそのインデックスを、なければ入るべきインデックスを返す。

## 私の解答

- 解法：二分探索

前問を活かし、`left = mid + 1` にしたかったが、解が `mid+1` にある場合面倒なので、if 文を多用してしまった。