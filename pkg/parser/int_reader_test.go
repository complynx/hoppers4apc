package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestParsers(t *testing.T) {
	suite.Run(t, new(parserTestSuite))
}

type parserTestSuite struct {
	suite.Suite
}

func (s *parserTestSuite) TestIntReaderOne() {
	s.Run("test fail reading", func() {
		r := strings.NewReader("")
		b := []byte{0}
		r.Read(b) // reach EOF

		reader := NewIntReader(r)
		_, err := reader.GetOne()
		s.EqualError(err, "EOF")
	})
	s.Run("test fail converting", func() {
		r := strings.NewReader("not_a_number\n")

		reader := NewIntReader(r)
		_, err := reader.GetOne()
		s.EqualError(err, "field 1 conversion error: strconv.Atoi: parsing \"not_a_number\": invalid syntax")
	})
	s.Run("test fail too few", func() {
		r := strings.NewReader("\n")

		reader := NewIntReader(r)
		_, err := reader.GetOne()
		s.EqualError(err, "expected one integer in the line")
	})
	s.Run("test fail too many", func() {
		r := strings.NewReader("42 28\n")

		reader := NewIntReader(r)
		_, err := reader.GetOne()
		s.EqualError(err, "expected one integer in the line")
	})
	s.Run("test ok", func() {
		r := strings.NewReader("42\n")

		reader := NewIntReader(r)
		num, err := reader.GetOne()
		s.NoError(err)
		s.Equal(42, num)
	})
	s.Run("dirty number ok", func() {
		r := strings.NewReader("\t42\r \n")

		reader := NewIntReader(r)
		num, err := reader.GetOne()
		s.NoError(err)
		s.Equal(42, num)
	})
	s.Run("double call", func() {
		r := strings.NewReader("42\n28\n")

		reader := NewIntReader(r)
		num, err := reader.GetOne()
		s.NoError(err)
		s.Equal(42, num)
		num, err = reader.GetOne()
		s.NoError(err)
		s.Equal(28, num)
	})
}

func (s *parserTestSuite) TestIntReaderTwo() {
	s.Run("test fail reading", func() {
		r := strings.NewReader("")
		b := []byte{0}
		r.Read(b) // reach EOF

		reader := NewIntReader(r)
		_, _, err := reader.GetCouple()
		s.EqualError(err, "EOF")
	})
	s.Run("test fail converting 1st", func() {
		r := strings.NewReader("not_a_number 12\n")

		reader := NewIntReader(r)
		_, _, err := reader.GetCouple()
		s.EqualError(err, "field 1 conversion error: strconv.Atoi: parsing \"not_a_number\": invalid syntax")
	})
	s.Run("test fail converting 2nd", func() {
		r := strings.NewReader("12 not_a_number\n")

		reader := NewIntReader(r)
		_, _, err := reader.GetCouple()
		s.EqualError(err, "field 2 conversion error: strconv.Atoi: parsing \"not_a_number\": invalid syntax")
	})
	s.Run("test fail too few", func() {
		r := strings.NewReader(" 2 \n")

		reader := NewIntReader(r)
		_, _, err := reader.GetCouple()
		s.EqualError(err, "expected two integers in the line")
	})
	s.Run("test fail too many", func() {
		r := strings.NewReader("42 28 12\n")

		reader := NewIntReader(r)
		_, _, err := reader.GetCouple()
		s.EqualError(err, "expected two integers in the line")
	})
	s.Run("test ok", func() {
		r := strings.NewReader("42 28\n")

		reader := NewIntReader(r)
		num, n2, err := reader.GetCouple()
		s.NoError(err)
		s.Equal(42, num)
		s.Equal(28, n2)
	})
	s.Run("dirty number ok", func() {
		r := strings.NewReader("\t42\r 28 \n")

		reader := NewIntReader(r)
		num, n2, err := reader.GetCouple()
		s.NoError(err)
		s.Equal(42, num)
		s.Equal(28, n2)
	})
	s.Run("double call", func() {
		r := strings.NewReader("42 12\n28 34\n")

		reader := NewIntReader(r)
		num, n2, err := reader.GetCouple()
		s.NoError(err)
		s.Equal(42, num)
		s.Equal(12, n2)
		num, n2, err = reader.GetCouple()
		s.NoError(err)
		s.Equal(28, num)
		s.Equal(34, n2)
	})
}

func (s *parserTestSuite) TestIntReaderFour() {
	s.Run("test fail reading", func() {
		r := strings.NewReader("")
		b := []byte{0}
		r.Read(b) // reach EOF

		reader := NewIntReader(r)
		_, _, _, _, err := reader.GetFour()
		s.EqualError(err, "EOF")
	})
	s.Run("test fail converting 1st", func() {
		r := strings.NewReader("not_a_number 12 12 23\n")

		reader := NewIntReader(r)
		_, _, _, _, err := reader.GetFour()
		s.EqualError(err, "field 1 conversion error: strconv.Atoi: parsing \"not_a_number\": invalid syntax")
	})
	s.Run("test fail converting 2nd", func() {
		r := strings.NewReader("12 not_a_number 33 31\n")

		reader := NewIntReader(r)
		_, _, _, _, err := reader.GetFour()
		s.EqualError(err, "field 2 conversion error: strconv.Atoi: parsing \"not_a_number\": invalid syntax")
	})
	s.Run("test fail converting 3nd", func() {
		r := strings.NewReader("12 33 not_a_number 31\n")

		reader := NewIntReader(r)
		_, _, _, _, err := reader.GetFour()
		s.EqualError(err, "field 3 conversion error: strconv.Atoi: parsing \"not_a_number\": invalid syntax")
	})
	s.Run("test fail converting 4th", func() {
		r := strings.NewReader("12 33 31 not_a_number\n")

		reader := NewIntReader(r)
		_, _, _, _, err := reader.GetFour()
		s.EqualError(err, "field 4 conversion error: strconv.Atoi: parsing \"not_a_number\": invalid syntax")
	})
	s.Run("test fail too few", func() {
		r := strings.NewReader(" 2 \n")

		reader := NewIntReader(r)
		_, _, _, _, err := reader.GetFour()
		s.EqualError(err, "expected four integers in the line")
	})
	s.Run("test fail too many", func() {
		r := strings.NewReader("42 28 12 23 32\n")

		reader := NewIntReader(r)
		_, _, _, _, err := reader.GetFour()
		s.EqualError(err, "expected four integers in the line")
	})
	s.Run("test ok", func() {
		r := strings.NewReader("42 28 33 41\n")

		reader := NewIntReader(r)
		num, n2, n3, n4, err := reader.GetFour()
		s.NoError(err)
		s.Equal(42, num)
		s.Equal(28, n2)
		s.Equal(33, n3)
		s.Equal(41, n4)
	})
	s.Run("dirty number ok", func() {
		r := strings.NewReader("\t42\r 28 33 41 \n")

		reader := NewIntReader(r)
		num, n2, n3, n4, err := reader.GetFour()
		s.NoError(err)
		s.Equal(42, num)
		s.Equal(28, n2)
		s.Equal(33, n3)
		s.Equal(41, n4)
	})
	s.Run("double call", func() {
		r := strings.NewReader("42 12 33 41\n28 34 65 234\n")

		reader := NewIntReader(r)
		num, n2, n3, n4, err := reader.GetFour()
		s.NoError(err)
		s.Equal(42, num)
		s.Equal(12, n2)
		s.Equal(33, n3)
		s.Equal(41, n4)
		num, n2, n3, n4, err = reader.GetFour()
		s.NoError(err)
		s.Equal(28, num)
		s.Equal(34, n2)
		s.Equal(65, n3)
		s.Equal(234, n4)
	})
}

func (s *parserTestSuite) TestIntReaderMultiple() {
	s.Run("ok", func() {
		r := strings.NewReader("42 12 33 41\n28 4\n3\n52 44\n")

		reader := NewIntReader(r)
		num, n2, n3, n4, err := reader.GetFour()
		s.NoError(err)
		s.Equal(42, num)
		s.Equal(12, n2)
		s.Equal(33, n3)
		s.Equal(41, n4)
		num, n2, err = reader.GetCouple()
		s.NoError(err)
		s.Equal(28, num)
		s.Equal(4, n2)
		num, err = reader.GetOne()
		s.NoError(err)
		s.Equal(3, num)
		num, n2, err = reader.GetCouple()
		s.NoError(err)
		s.Equal(52, num)
		s.Equal(44, n2)
	})
	s.Run("fail", func() {
		r := strings.NewReader("42 12 33 41\n28 4\n3\n52 44\n")

		reader := NewIntReader(r)
		num, n2, n3, n4, err := reader.GetFour()
		s.NoError(err)
		s.Equal(42, num)
		s.Equal(12, n2)
		s.Equal(33, n3)
		s.Equal(41, n4)
		num, n2, err = reader.GetCouple()
		s.NoError(err)
		s.Equal(28, num)
		s.Equal(4, n2)
		_, _, err = reader.GetCouple()
		s.EqualError(err, "expected two integers in the line")
	})
}
