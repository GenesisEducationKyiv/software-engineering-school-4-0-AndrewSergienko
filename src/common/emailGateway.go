package common

type EmailSender interface {
	Send(target string, rate float32)
}
