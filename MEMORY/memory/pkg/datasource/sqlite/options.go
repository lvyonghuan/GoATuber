package sqlite

type options struct {
	path string
}

// CallOptions provides a way to set options.
type CallOptions struct {
	applyFunc func(*options)
}

func applyCallOptions(callOptions []CallOptions, defaultOptions ...options) *options {
	o := new(options)
	if len(defaultOptions) > 0 {
		*o = defaultOptions[0]
	}

	for _, callOption := range callOptions {
		callOption.applyFunc(o)
	}
	return o
}

// WithPath sets the path option.
func WithPath(path string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.path = path
		},
	}
}
