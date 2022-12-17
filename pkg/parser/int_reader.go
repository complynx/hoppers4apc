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

// NewIntReader creates a buffered reader and the int reader itself.
// It takes an io.Reader as an argument and returns a pkg.IntReader.
func NewIntReader(r io.Reader) pkg.IntReader {
	return &intParser{
		reader: *bufio.NewReader(r),
	}
}

// getFields is a supplementary function that gets some number of integers on a line separated by some spaces.
// It reads a line from the input stream, splits it into word chunks, and parses each chunk to int.
// It returns an array of integers and an error if there was any problem with the conversion.
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

// GetOne reads a line from the input stream and expects one integer there.
// It returns the integer and an error if there was a problem with the
// conversion or if there was not exactly one integer in the line.
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

// GetCouple reads a line from the input stream and expects two integers there.
// It returns the two integers and an error if there was a problem with the
// conversion or if there were not exactly two integers in the line.
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

// GetFour reads a line from the input stream and expects four integers there.
// It returns the four integers and an error if there was a problem with the
// conversion or if there were not exactly four integers in the line.
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
