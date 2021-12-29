package acom2000

import (
	"errors"
	"fmt"
	"strconv"
)

func dec2Ascii(dec int) (string, error) {

	if dec < 0 {
		return "", errors.New("decimal must not be negative")
	}

	hex := fmt.Sprintf("%x", dec)
	ascii := ""
	for _, digit := range hex {
		switch digit {
		case 'a':
			ascii = ascii + ":"
		case 'b':
			ascii = ascii + ";"
		case 'c':
			ascii = ascii + "<"
		case 'd':
			ascii = ascii + "="
		case 'e':
			ascii = ascii + ">"
		case 'f':
			ascii = ascii + "?"
		default:
			ascii = ascii + string(digit)
		}
	}
	return ascii, nil
}

func ascii2Dec(s string) (int, error) {

	hex := ""

	for _, digit := range s {
		switch string(digit) {
		case ":":
			hex = hex + "a"
		case ";":
			hex = hex + "b"
		case "<":
			hex = hex + "c"
		case "=":
			hex = hex + "d"
		case ">":
			hex = hex + "e"
		case "?":
			hex = hex + "f"
		case "0":
			hex = hex + "0"
		case "1":
			hex = hex + "1"
		case "2":
			hex = hex + "2"
		case "3":
			hex = hex + "3"
		case "4":
			hex = hex + "4"
		case "5":
			hex = hex + "5"
		case "6":
			hex = hex + "6"
		case "7":
			hex = hex + "7"
		case "8":
			hex = hex + "8"
		case "9":
			hex = hex + "9"
		default:
			return 0, errors.New("illegal character")
		}

	}
	dec, err := strconv.ParseInt(hex, 16, 0)
	if err != nil {
		return 0, err
	}

	return int(dec), nil
}
