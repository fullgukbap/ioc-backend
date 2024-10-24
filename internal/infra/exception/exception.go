package exception

type Exception struct {
	Type    string
	Status  int
	Message string
	Data    any

	Err error
}

type ExceptionOption func(*Exception)

func WithData(data any) ExceptionOption {
	return func(e *Exception) {
		e.Data = data
	}
}

func (e Exception) Error() string {
	return e.Err.Error()
}

type Map map[string]any

func New(err error, _type string, status int, message string, opts ...ExceptionOption) *Exception {
	exception := &Exception{
		Err:     err,
		Type:    _type,
		Status:  status,
		Message: message,
	}

	// 옵션 적용
	for _, opt := range opts {
		opt(exception)
	}

	return exception
}

func NewWithoutErr(_type string, status int, message string, opts ...ExceptionOption) *Exception {
	exception := &Exception{
		Type:    _type,
		Status:  status,
		Message: message,
	}

	// 옵션 적용
	for _, opt := range opts {
		opt(exception)
	}

	return exception
}
