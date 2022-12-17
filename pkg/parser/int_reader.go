package parser

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/complynx/hoppers4apc/pkg"
)

// this is a parser for integer input lines
type intParser struct {
	reader bufio.Reader
}

// creates a buffered reader and the int reader itself
func NewIntReader(r io.Reader) pkg.IntReader {
	return &intParser{
		reader: *bufio.NewReader(r),
	}
}

// supplementary function to get some number of integers on a line separated by some spaces
func (p *intParser) getFields() ([]int, error) {
	// get the line
	line, err := p.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	// split it into word chunks
	fields := strings.Fields(line)
	// create output array of same size
	ret := make([]int, len(fields))
	// parse each chunk to int and save to array
	for i := range fields {
		num, err := strconv.Atoi(fields[i])
		if err != nil {
			return nil, fmt.Errorf("field %d conversion error: %w", i+1, err)
		}
		ret[i] = num
	}
	return ret, nil
}

// reads line from input stream and expects one integer there
func (p *intParser) GetOne() (int, error) {
	line, err := p.getFields()
	if err != nil {
		return 0, err
	}
	if len(line) != 1 {
		return 0, errors.New("expected one integer in the line")
	}
	return line[0], nil
}

// reads line from input stream and expects two integers there
func (p *intParser) GetCouple() (int, int, error) {
	line, err := p.getFields()
	if err != nil {
		return 0, 0, err
	}
	if len(line) != 2 {
		return 0, 0, errors.New("expected two integers in the line")
	}
	return line[0], line[1], nil
}

// reads line from input stream and expects four integers there
func (p *intParser) GetFour() (int, int, int, int, error) {
	line, err := p.getFields()
	if err != nil {
		return 0, 0, 0, 0, err
	}
	if len(line) != 4 {
		return 0, 0, 0, 0, errors.New("expected four integers in the line")
	}
	return line[0], line[1], line[2], line[3], nil
}
