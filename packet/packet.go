package packet

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Payload struct {
	Event string
	Body  []byte
}

const delim = '\n'

// Decode the event from a packet
func DecodeEvent(r io.Reader, length int, evt *string) (n int, e error) {
	br := bufio.NewReader(r)
	header, e := br.ReadString(delim)
	if e != nil && e != io.EOF {
		fmt.Println(e)
		return n, e
	}
	headerParts := strings.SplitN(header, ":", 2)
	if len(headerParts) < 2 {
		return n, errors.New("Malformed payload header.")
	}

	*evt = strings.TrimSpace(headerParts[1])
	return n, e
}

// Encode a full packet
func Encode(w io.Writer, p *Payload) (n int, err error) {
	var i int
	i, err = io.WriteString(w, fmt.Sprintf("event: %s\n", p.Event))
	n += i
	if err != nil {
		return n, err
	}
	// Encode body
	if p.Body != nil {
		i, err = w.Write(p.Body)
		n += i
		if err != nil {
			return n, err
		}
	}
	return n, err
}

// Encode a packet with no body
func EncodeEvent(w io.Writer, event string) (n int, err error) {
	return Encode(w, &Payload{
		Event: event})
}
