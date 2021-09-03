package foo

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrsmsn/go-sandbox/mockgen/mock_foo"
)

func TestFoo(t *testing.T) {
	// コントローラーの生成
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// モックの生成
	m := mock_foo.NewMockFoo(ctrl)

	// 対象メソッドと引数と返り値を設定する
	m.EXPECT().Bar(99).Return(1000)

	// モックの呼び出し
	m.Bar(99)
}
