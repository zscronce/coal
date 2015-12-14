package coal

// given the integer value of an attribute, computes a new value for said attribute
type mutation interface {
	apply(int) int
}

type additiveMutation struct {
	delta int
}

func (this *additiveMutation) apply(attribute int) int {
	return attribute + this.delta
}
