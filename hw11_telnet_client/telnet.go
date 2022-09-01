package main

import (
	"context"
	"errors"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	Close() error
	Send() error
	Receive() error
}

type Telnet struct {
	address    string
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	conn       net.Conn
	cancelFunc context.CancelFunc
}

func NewTelnetClient(
	address string,
	timeout time.Duration,
	in io.ReadCloser,
	out io.Writer,
	cancelFunc context.CancelFunc,
) TelnetClient {
	return &Telnet{
		address:    address,
		timeout:    timeout,
		in:         in,
		out:        out,
		cancelFunc: cancelFunc,
	}
}

func (t *Telnet) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err == nil {
		t.conn = conn
	}
	return err
}

func (t *Telnet) Close() error {
	if t.conn != nil {
		if err := t.conn.Close(); err != nil {
			t.cancelFunc()
			return err
		}
	}
	return nil
}

func (t *Telnet) Send() error {
	if _, err := io.Copy(t.conn, t.in); err != nil && !errors.Is(err, io.EOF) {
		t.cancelFunc()
		return err
	}

	return nil
}

func (t *Telnet) Receive() error {
	if _, err := io.Copy(t.out, t.conn); err != nil && !errors.Is(err, io.EOF) {
		t.cancelFunc()
		return err
	}

	return nil
}
