package Problem0006

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	res := bytes.Buffer{}
	// p pace
	p := numRows*2 - 2

	// process the first line
	for i := 0; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	// process the middle row
	for r := 1; r <= numRows-2; r++ {
		// add the first character of line r
		res.WriteByte(s[r])

		for k := p; k-r < len(s); k += p {
			res.WriteByte(s[k-r])
			if k+r < len(s) {
				res.WriteByte(s[k+r])
			}
		}
	}

	// process the last line
	for i := numRows - 1; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	return res.String()
}