package enotify

// RichSender is the interface type for sending rich notifications.
type RichSender interface {
	Send(title string, body string, link string) bool
}
