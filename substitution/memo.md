# 代入したメソッドの引数にstructを指定することで、そのstructのメソッドとして実行できる。

## 用途
`unexported`な関数をテスト用に`exported`にしたり。
でも次の方法のが一般的らしい
```go
func (f *Strunct) ExportedFunc(arg1, arg2 string) {
  f.unexportedFunc(arg1, arg2)
}
```

## その他
ちゃんと調べてないけど、多分レシーバーとかメソッドあたりの言語仕様をしっかり調べれば出そう。
