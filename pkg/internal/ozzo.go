package internal

type validatable interface {
	Validate() error
}

func Validate(v validatable) error {
	return v.Validate()
}
