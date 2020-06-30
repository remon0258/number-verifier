package providers

// Providers is the interface where all providers need to confirm to.
type Providers interface {
	GetProvider() Provider
	GetNumbers() ([]string, error)
	GetMessages(number string) ([]string, error)
}

// Provider is a struct which contains some properties about a provider.
type Provider struct {
	Name string
	BaseURL string
}