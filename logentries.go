package logentries

import (
	"crypto/tls"
	"io"
	"net"
)

type logEntriesWriter struct {
	token         string
	outputStream  io.Writer
	wrappedStream io.Writer
}

// NewLogEntriesWriter creates a new io.Writer which can then be wrapped in a Logger.
// You do not have to use it within a Logger, but that is the intended method.
// You must pass in the token and secure parameters.
// You can optionally pass in wrappedStream which will, in addition to writting to Log Entries, also have data written to it.
// This allows you to log to the console while still logging to Log Entries.
// If you do not wish to use wrappedStream, pass in nil.
func NewLogEntriesWriter(token string, secure bool, wrappedStream io.Writer) (io.Writer, error) {
	var outputStream net.Conn
	var err error
	if secure {
		config := tls.Config{}
		outputStream, err = tls.Dial("tcp", "api.logentries.com:20000", &config)
		if err != nil {
			return nil, err
		}
	} else {
		outputStream, err = net.Dial("tcp", "data.logentries.com:80")
		if err != nil {
			return nil, err
		}
	}
	return &logEntriesWriter{token, outputStream, wrappedStream}, nil
}

func (l *logEntriesWriter) Write(p []byte) (int, error) {
	_, err := l.outputStream.Write([]byte(l.token))
	if err != nil {
		return 0, err
	}

	count, err := l.outputStream.Write(p)
	if err != nil {
		return count, err
	}
	if l.wrappedStream != nil {
		count, err := l.wrappedStream.Write(p)
		if err != nil {
			return count, err
		}
	}
	return count, nil
}
