package notifier

type Notifier interface {
	Push(*Request) (string, error)
}

type Request struct {
	Title    string
	Contents string
}
