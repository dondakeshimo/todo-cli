package notifier

type Notifier interface {
	Push(*Request) error
}

type Request struct {
	Title string
	Contents string
}
