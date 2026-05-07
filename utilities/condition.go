package utilities

func Ternary[T any](condition bool, trueValue T, elseValue T) T {
	if condition {
		return trueValue
	}
	return elseValue
}

func TernaryF[T any](condition bool, trueValue func() T, elseValue func() T) T {
	if condition {
		return trueValue()
	}
	return elseValue()
}

func LogicalOrInt[T Signed](value T, elValue T) T {
	if value == 0 {
		return elValue
	}
	return value
}

func LogicalOrFloat[T Float](value T, elValue T) T {
	if value == 0 {
		return elValue
	}
	return value
}

func LogicalOrString[T ~string](value T, elValue T) T {
	if len(value) == 0 {
		return elValue
	}
	return value
}

func LogicalError(err error, condition bool, errCondition error) error {
	if err != nil {
		return err
	}
	if condition && errCondition != nil {
		return errCondition
	}
	return nil
}
