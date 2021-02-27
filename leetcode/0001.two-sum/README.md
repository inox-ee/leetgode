# Two Sum

## 問題

与えられた int 配列から、その和が target となるようなある 2 つの数の組のインデックスを求める。

引数：nums []int, target int

## 私の解答

- 解法：ブルートフォース
- 計算量：O(n^2)

nums[i] (i := 0...len(nums)) に対し、nums[j] (j := i+1...len(nums)) の和を求め、target と一致するか調べる。

## 解答１

- 解法：Two-pass Hash Table
- 計算量：O(n)

ハッシュテーブルを用いる。
ハッシュテーブルの key に nums の値を、value にその値のインデックスを登録することで、target - nums[i] の値が存在するかしないかを O(1)で求められる。

## 解答２

- 解法；One-pass Hash Table
- 計算量：O(n)

全て登録せずとも、代入しながら判定すればよい。
