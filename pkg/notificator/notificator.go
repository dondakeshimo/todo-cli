package notificator

type Notificator interface {
	Push(*Request) error
}

type Request struct {
	Title string
	Contents string
}
