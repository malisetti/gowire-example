package hello

type Transformer interface {
	Transform(string) string
}

type (
	NewLine struct{}
	Exact   struct{}
	Zero    struct{}
)

func NewDefaultTransformer() Transformer {
	return NewLine{}
}

func (NewLine) Transform(s string) string {
	return s + "\n"
}

func (Exact) Transform(s string) string {
	return s
}

func (Zero) Transform(s string) string {
	return ""
}
