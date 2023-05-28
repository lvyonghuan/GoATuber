package boltdb

import "os"

type options struct {
	path   string
	bucket string
	mode   os.FileMode
}
type CallOptions struct {
	applyFunc func(o *options)
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

func WithPath(path string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.path = path
		},
	}
}

func WithBucket(bucket string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.bucket = bucket
		},
	}
}
