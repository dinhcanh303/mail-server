package mail

type EmailSender interface {
	Configure(...Option) EmailSender
	SendEmail(subject, content string, to, cc, bcc []string, attachFiles []string) error
}
