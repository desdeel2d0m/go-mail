package mail

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"strings"
	"testing"
	"time"
)

// brokenWriter implements a broken writer for io.Writer testing
type brokenWriter struct {
	io.Writer
}

// Write implements the io.Writer interface but intentionally returns an error at
// any time
func (bw *brokenWriter) Write([]byte) (int, error) {
	return 0, fmt.Errorf("intentionally failed")
}

// TestMsgWriter_Write tests the WriteTo() method of the msgWriter
func TestMsgWriter_Write(t *testing.T) {
	bw := &brokenWriter{}
	mw := &msgWriter{w: bw, c: CharsetUTF8, en: mime.QEncoding}
	_, err := mw.Write([]byte("test"))
	if err == nil {
		t.Errorf("msgWriter WriteTo() with brokenWriter should fail, but didn't")
	}

	// Also test the part when a previous error happened
	mw.err = fmt.Errorf("broken")
	_, err = mw.Write([]byte("test"))
	if err == nil {
		t.Errorf("msgWriter WriteTo() with brokenWriter should fail, but didn't")
	}
}

// TestMsgWriter_writeMsg tests the writeMsg method of the msgWriter
func TestMsgWriter_writeMsg(t *testing.T) {
	m := NewMsg()
	_ = m.From(`"Toni Tester" <test@example.com>`)
	_ = m.To(`"Toni Receiver" <receiver@example.com>`)
	m.Subject("This is a subject")
	m.SetBulk()
	now := time.Now()
	m.SetDateWithValue(now)
	m.SetMessageIDWithValue("message@id.com")
	m.SetBodyString(TypeTextPlain, "This is the body")
	buf := bytes.Buffer{}
	mw := &msgWriter{w: &buf, c: CharsetUTF8, en: mime.QEncoding}
	mw.writeMsg(m)
	ms := buf.String()

	var ea []string
	if !strings.Contains(ms, `MIME-Version: 1.0`) {
		ea = append(ea, "MIME-Version")
	}
	if !strings.Contains(ms, fmt.Sprintf("Date: %s", now.Format(time.RFC1123Z))) {
		ea = append(ea, "Date")
	}
	if !strings.Contains(ms, `Message-ID: <message@id.com>`) {
		ea = append(ea, "Message-ID")
	}
	if !strings.Contains(ms, `Precedence: bulk`) {
		ea = append(ea, "Precedence")
	}
	if !strings.Contains(ms, `Subject: This is a subject`) {
		ea = append(ea, "Subject")
	}
	if !strings.Contains(ms, `User-Agent: go-mail v`) {
		ea = append(ea, "User-Agent")
	}
	if !strings.Contains(ms, `X-Mailer: go-mail v`) {
		ea = append(ea, "X-Mailer")
	}
	if !strings.Contains(ms, `From: "Toni Tester" <test@example.com>`) {
		ea = append(ea, "From")
	}
	if !strings.Contains(ms, `To: "Toni Receiver" <receiver@example.com>`) {
		ea = append(ea, "To")
	}
	if !strings.Contains(ms, `Content-Type: text/plain; charset=UTF-8`) {
		ea = append(ea, "Content-Type")
	}
	if !strings.Contains(ms, `Content-Transfer-Encoding: quoted-printable`) {
		ea = append(ea, "Content-Transfer-Encoding")
	}
	if !strings.Contains(ms, "\r\n\r\nThis is the body") {
		ea = append(ea, "Message body")
	}
	if len(ea) > 0 {
		em := "writeMsg() failed. The following errors occurred:\n"
		for e := range ea {
			em += fmt.Sprintf("* incorrect %q field", ea[e])
		}
		em += fmt.Sprintf("\n\nFull message:\n%s", ms)
		t.Errorf(em)
	}
}
