package notificator

type Notificator interface {
	Push() (Response, error)
}

type Response struct {
	string
}

type Scheduler interface {
	Register() ()
}
