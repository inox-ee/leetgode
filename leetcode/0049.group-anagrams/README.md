# アナグラムのグループ化

アナグラム文字列を、使っている文字ごとにグループ分けする。

## 私の解答

- 解法：各単語をソートし、ハッシュマップで存在確認

各単語をアルファベット順にソートするので結構計算量が多い。
~~あとマップ(`map[string][]string`) の value にはソート前の単語を入れたが、メモリが無駄なので、与配列のインデックスにすべきだったかも~~。-> なんかゴミみたいな結果になった。

## 解答例１

- 各アルファベットの使用個数をカウント

アルファベットは高々 26 個しかないことに着目。

## 解答例２

- 各アルファベットに素数を割り振り、与えられた単語を素数の積で表現

天才だろ…

## 結果

    Runtime: 24 ms, faster than 78.66% of Go online submissions for Group Anagrams.
    Memory Usage: 8.1 MB, less than 37.80% of Go online submissions for Group Anagrams.
