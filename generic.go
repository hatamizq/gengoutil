package gengoutil

// get zero value of type
func ZeroOf[T any]() (_ T) { return }

// get pointer value of type
func Ptr[T any](v T) *T {
	return &v
}

// get nil if the value is empty
func NilEmpty[T comparable](v T) *T {
	if v == ZeroOf[T]() {
		return nil
	}
	return Ptr(v)
}

// get empty if the value is nil
func EmptyNil[T comparable](v *T) T {
	if v == nil {
		return ZeroOf[T]()
	}
	return *v
}

// check whether the value is empty
func IsEmpty[T comparable](v T) bool {
	return v == ZeroOf[T]()
}

// check whether the value is not empty
func IsNotEmpty[T comparable](v T) bool {
	return v != ZeroOf[T]()
}

// coalesce value if empty
func Coalesce[T comparable](v T, def T) T {
	if NilEmpty(v) == nil {
		return def
	}
	return v
}

// coalesce ptr value if empty
func CoalescePtr[T comparable](v *T, def *T) *T {
	if v == nil {
		return def
	}
	return v
}
