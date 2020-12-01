package notifier

// Notifier is a interface that notify some string.
type Notifier interface {
	Push(*Request) (string, error)
}

// Request is a struct that is passed to Notifier.
type Request struct {
	Title    string
	Contents string
	Answer   []string
}
