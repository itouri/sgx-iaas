package driver

type Driver interface {
	Send() // reply message
	SendNotification()
	Listen()
	ListenNotification()
	CleanUp()
}
