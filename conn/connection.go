package conn

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

// New returns a net.Conn over TCP, optionally secured with TLS.
func New(address string, useTLS bool) (net.Conn, error) {
	if useTLS {
		conn, err := tls.Dial("tcp", address, &tls.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect (TLS): %w", err)
		}
		return conn, nil
	}

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect (non-TLS): %w", err)
	}
	return conn, nil
}

// SendCommand serializes any struct as JSON and writes it to the connection.
func SendCommand(w io.Writer, cmd any) error {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(cmd); err != nil {
		return fmt.Errorf("failed to send command: %w", err)
	}
	return nil
}
