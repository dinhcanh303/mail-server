package mail

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendMail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	senderMail := NewEmailSender()
	subject := "Email test"
	content := `<h1>Hello world</h1>
	<p>This is a test message from <a href="https://github.com/dinhcanh303">Foden Ngo</a></p>`

	to := []string{"dinhcanhng303@gmail.com"}
	// attachFiles := []string{"../../README.md"}
	err := senderMail.SendEmail(subject, content, to, nil, nil, nil)
	require.NoError(t, err)
}
