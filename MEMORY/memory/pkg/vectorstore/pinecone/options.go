package pc

type options struct {
	apiKey      string
	indexName   string
	namespace   string
	projectName string
	environment string
}

// CallOptions provides a way to set optional parameters to various methods
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

// WithApiKey sets the API key for the Pinecone client
func WithApiKey(apiKey string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.apiKey = apiKey
		},
	}
}

// WithIndexName sets the index name for the Pinecone client
func WithIndexName(indexName string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.indexName = indexName
		},
	}
}

// WithNamespace sets the namespace for the Pinecone client
// asllm is the default namespace
func WithNamespace(namespace string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.namespace = namespace
		},
	}
}

// WithProjectName sets the project name for the Pinecone client
func WithProjectName(projectName string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.projectName = projectName
		},
	}
}

// WithEnvironment sets the environment for the Pinecone client
// aka the region
func WithEnvironment(environment string) CallOptions {
	return CallOptions{
		applyFunc: func(o *options) {
			o.environment = environment
		},
	}
}
