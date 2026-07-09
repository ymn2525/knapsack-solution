# knapsack-solution
知能情報演習IV 課題

## 取り組んだこと
- 課題1: ナップサック問題総当り
- 課題2: ナップサック問題 動的計画法

## 実行方法
`
go run $(ls $(ls *.go | grep -v _test.go)
`

## benchmark計測
`
go test -bench . -benchmem
`

## git/githubの設定完了の証拠
ローカルリポジトリで編集してリモートリポジトリにpush成功
