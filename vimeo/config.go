package vimeo

// Config provides a way to configure the Client depending on your needs.
type Config struct {
	// Uploader
	Uploader Uploader
}

// DefaultConfig return the default Client configuration.
func DefaultConfig() *Config {
	return &Config{
		Uploader: nil,
	}
}
