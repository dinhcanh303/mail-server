package mail

type Option func(*emailSender)

func Host(host string) Option {
	return func(e *emailSender) {
		e.host = host
	}
}
func Port(port string) Option {
	return func(e *emailSender) {
		e.port = port
	}
}

func Username(username string) Option {
	return func(e *emailSender) {
		e.username = username
	}
}

func Password(password string) Option {
	return func(e *emailSender) {
		e.password = password
	}
}

func MaxConnections(maxConnections int) Option {
	return func(e *emailSender) {
		e.maxConnections = maxConnections
	}
}
func IDLETimeout(idleTimeout int) Option {
	return func(e *emailSender) {
		e.idleTimeout = idleTimeout
	}
}
func WaitTimeout(waitTimeout int) Option {
	return func(e *emailSender) {
		e.waitTimeout = waitTimeout
	}
}

func Retries(retries int) Option {
	return func(e *emailSender) {
		e.retries = retries
	}
}
func SkipTLS(skipTls bool) Option {
	return func(e *emailSender) {
		e.skipTls = skipTls
	}
}
