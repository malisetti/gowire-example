package hello

import (
	"github.com/google/wire"
)

type Transformer interface {
	Transform(string) string
}

type (
	NewLine struct{}
	Exact   struct{}
	Zero    struct{}
)

type TransformType int

const (
	ZeroTransform TransformType = iota
	NewLineTransform
	ExactTransform
)

type TransformerProviderConfig interface {
	GetTransformerProviderType() TransformType
}

func NewTransform(cfg TransformerProviderConfig) Transformer {
	switch cfg.GetTransformerProviderType() {
	case ZeroTransform:
		return NewZeroTransformer()
	case NewLineTransform:
		return NewLineTransformer()
	case ExactTransform:
		return NewExactTransformer()
	default:
		panic("invalid transform")
	}
}

func NewLineTransformer() *NewLine {
	return &NewLine{}
}

func NewExactTransformer() *Exact {
	return &Exact{}
}

func NewZeroTransformer() *Zero {
	return &Zero{}
}

var TransformSet = wire.NewSet(
	NewTransform,
)

func (*NewLine) Transform(s string) string {
	return s + "\n"
}

func (*Exact) Transform(s string) string {
	return s
}

func (*Zero) Transform(s string) string {
	return ""
}
