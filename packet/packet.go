package packet

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/gobwas/ws"
)

type Payload struct {
	Event string
	Body  []byte
}

const delim = '\n'

func Decode(r io.Reader, wsHeader ws.Header) (p *Payload, err error) {
	p = new(Payload)
	// Parse event type
	br := bufio.NewReader(r)
	// Payload header
	header, err := br.ReadString(delim)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return nil, err
	}
	headerParts := strings.SplitN(header, ":", 2)
	if len(headerParts) < 2 {
		return nil, errors.New("Malformed payload header.")
	}
	// Set event
	headerParts[1] = headerParts[1][0 : len(headerParts[1])-1] // Remove trailing newline
	p.Event = strings.Trim(headerParts[1], " ")
	// Read body
	p.Body = make([]byte, wsHeader.Length-int64(len(header)))
	_, err = io.ReadFull(r, p.Body)

	return p, err
}

func Encode(w io.Writer, p *Payload) (n int, err error) {
	var i int
	i, err = io.WriteString(w, fmt.Sprintf("event: %s\n", p.Event))
	n += i
	if err != nil {
		return n, err
	}
	if p.Body != nil {
		switch p.Body {

		}
	}
	return n, err
}

func EncodeEvent(w io.Writer, event string) (n int, err error) {
	return Encode(w, &Payload{
		Event: event})
}
