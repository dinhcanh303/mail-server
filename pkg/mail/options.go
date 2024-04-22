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

func AuthProtocol(authProtocol string) Option {
	return func(e *emailSender) {
		e.authProtocol = authProtocol
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

func FromName(fromName string) Option {
	return func(e *emailSender) {
		e.fromName = fromName
	}
}

func FromAddress(fromAddress string) Option {
	return func(e *emailSender) {
		e.fromAddress = fromAddress
	}
}

func MaxConnections(maxConnections int64) Option {
	return func(e *emailSender) {
		e.maxConnections = maxConnections
	}
}
func IDLETimeout(idleTimeout int64) Option {
	return func(e *emailSender) {
		e.idleTimeout = idleTimeout
	}
}
func WaitTimeout(waitTimeout int64) Option {
	return func(e *emailSender) {
		e.waitTimeout = waitTimeout
	}
}

func Retries(retries int64) Option {
	return func(e *emailSender) {
		e.retries = retries
	}
}
func TypeTLS(typeTls string) Option {
	return func(e *emailSender) {
		e.tlsType = typeTls
	}
}
