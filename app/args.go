package app

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/anihouse/bot/app/dstype"
)

// Args errors
var (
	ErrNotMatch = errors.New("data not math with template")
)

var (
	regexpInt = regexp.MustCompile(`^\S+`)
)

type args struct {
	Args string
}

func (a *args) Scan(dest ...interface{}) (err error) {
	var (
		lenArgs = len(a.Args)

		iDest = 0
		iArgs = skipSpaces(a.Args, 0)
	)

	if iArgs == -1 {
		return nil
	}

	for {
		if iArgs >= lenArgs || iDest >= len(dest) {
			break
		}

		iArgs = skipSpaces(a.Args, iArgs)
		if iArgs == -1 {
			return nil
		}

		switch t := dest[iDest].(type) {
		case *string:
			value := dest[iDest].(*string)
			*value, err = doString(a.Args, &iArgs)
			if err != nil {
				return
			}
			return nil
		case *int:
			value := dest[iDest].(*int)
			*value, err = doInt(a.Args, &iArgs)
			if err != nil {
				return
			}
		default:
			value, ok := t.(dstype.Scanneable)
			if !ok {
				return fmt.Errorf("type %T not scannable", t)
			}
			err = value.Scan(a.Args, &iArgs)
			if err != nil {
				return
			}
		}
		iDest++
		iArgs++
	}
	return nil
}

func (a *args) Scanf(tpl string, dest ...interface{}) (err error) {
	if count := strings.Count(tpl, "?"); count != len(dest) {
		return fmt.Errorf("count '?' [%d] and length 'dest' [%d] not equal", count, len(dest))
	}

	tpl = normalizeTpl(tpl)

	var (
		lenTpl  = len(tpl)
		lenArgs = len(a.Args)

		iDest = 0
		iTpl  = 0
		iArgs = skipSpaces(a.Args, 0)
	)

	if iArgs == -1 {
		return nil
	}

	for {
		if iTpl >= lenTpl || iArgs >= lenArgs || iDest >= len(dest) {
			break
		}

		iArgs = skipSpaces(a.Args, iArgs)
		if iArgs == -1 {
			return nil
		}

		if tpl[iTpl] == '?' {
			switch t := dest[iDest].(type) {
			case *string:
				value := dest[iDest].(*string)
				*value, err = doString(a.Args, &iArgs)
				if err != nil {
					return
				}
				return nil
			case *int:
				value := dest[iDest].(*int)
				*value, err = doInt(a.Args, &iArgs)
				if err != nil {
					return
				}
			default:
				value, ok := t.(dstype.Scanneable)
				if !ok {
					return fmt.Errorf("type %T not scannable", t)
				}
				err = value.Scan(a.Args, &iArgs)
				if err != nil {
					return
				}
			}
			iDest++
		} else {
			if tpl[iTpl] != a.Args[iArgs] {
				return ErrNotMatch
			}
		}
		iTpl++
		iArgs++
	}

	return nil
}

func normalizeTpl(tpl string) string {
	placeholders := strings.Split(tpl, "?")

	if len(placeholders) == 0 {
		return tpl
	}

	for i, p := range placeholders {
		placeholders[i] = strings.TrimSpace(p)
	}
	return strings.Join(placeholders, "?")
}

func skipSpaces(s string, i int) int {
	for i < len(s) {
		if s[i] != ' ' {
			return i
		}
		i++
	}
	return -1
}

func doInt(s string, i *int) (int, error) {
	loc := regexpInt.FindIndex([]byte(s[*i:]))
	number := s[*i : *i+loc[1]]

	value, err := strconv.Atoi(number)
	if err != nil {
		return 0, err
	}

	*i += loc[1]
	return value, nil
}

func doString(s string, i *int) (string, error) {
	return s[*i:], nil
}
