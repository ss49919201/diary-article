package dicontainer

import (
	"github.com/samber/do"
)

// コンポーネントの初期化処理の煩雑さ、New関数の引数の肥大化を防ぐために `samber/do` を使う

var injector = do.New()

type Provider[T any] func() (T, error)

func provider[T any](fn Provider[T]) do.Provider[T] {
	return func(in *do.Injector) (T, error) {
		return fn()
	}
}

func Provide[T any](fn Provider[T]) {
	do.Provide(injector, provider(fn))
}

func MustInvoke[T any]() T {
	return do.MustInvoke[T](injector)
}
