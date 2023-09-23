package hello

import "github.com/google/wire"

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
		return newZeroTransformer()
	case NewLineTransform:
		return newLineTransformer()
	default:
		return newExactTransformer()
	}
}

func newLineTransformer() *NewLine {
	return &NewLine{}
}

func newExactTransformer() *Exact {
	return &Exact{}
}

func newZeroTransformer() *Zero {
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
