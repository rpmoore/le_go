package logentries

import (
	"io"
	"net"
)

type logEntriesWriter struct {
	token        string
	outputStream io.Writer
}

func NewLogEntriesWriter(token string, secure bool) (io.Writer, error) {
	var outputStream net.Conn
	/*
		if secure {
			outputStream = tls.Dial("tcp", endpoint)
		}
	*/
	outputStream, err := net.Dial("tcp", "data.logentries.com:80")
	if err != nil {
		return nil, err
	}
	return &logEntriesWriter{token, outputStream}, nil
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
	return count, nil
}
