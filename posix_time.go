package posix_time

import (
	"errors"
	"fmt"
	"strings"
)

/*
pattern	description
%A	national representation of the full weekday name
%a	national representation of the abbreviated weekday
%B	national representation of the full month name
%b	national representation of the abbreviated month name
%C	(year / 100) as decimal number; single digits are preceded by a zero
%c	national representation of time and date
%D	equivalent to %m/%d/%y
%d	day of the month as a decimal number (01-31)
%e	the day of the month as a decimal number (1-31); single digits are preceded by a blank
%F	equivalent to %Y-%m-%d
%H	the hour (24-hour clock) as a decimal number (00-23)
%h	same as %b
%I	the hour (12-hour clock) as a decimal number (01-12)
%j	the day of the year as a decimal number (001-366)
%k	the hour (24-hour clock) as a decimal number (0-23); single digits are preceded by a blank
%l	the hour (12-hour clock) as a decimal number (1-12); single digits are preceded by a blank
%M	the minute as a decimal number (00-59)
%m	the month as a decimal number (01-12)
%n	a newline
%p	national representation of either "ante meridiem" (a.m.) or "post meridiem" (p.m.) as appropriate.
%R	equivalent to %H:%M
%r	equivalent to %I:%M:%S %p
%S	the second as a decimal number (00-60)
%T	equivalent to %H:%M:%S
%t	a tab
%U	the week number of the year (Sunday as the first day of the week) as a decimal number (00-53)
%u	the weekday (Monday as the first day of the week) as a decimal number (1-7)
%V	the week number of the year (Monday as the first day of the week) as a decimal number (01-53)
%v	equivalent to %e-%b-%Y
%W	the week number of the year (Monday as the first day of the week) as a decimal number (00-53)
%w	the weekday (Sunday as the first day of the week) as a decimal number (0-6)
%X	national representation of the time
%x	national representation of the date
%Y	the year with century as a decimal number
%y	the year without century as a decimal number (00-99)
%Z	the time zone name
%z	the time zone offset from UTC
%%	a '%'

Go:
These are predefined layouts for use in Time.Format and time.Parse. The reference time used in the layouts is the specific time:

Mon Jan 2 15:04:05 MST 2006
which is Unix time 1136239445. Since MST is GMT-0700, the reference time can be thought of as

01/02 03:04:05PM '06 -0700
To define your own format, write down what the reference time would look like formatted your way; see the values of constants like ANSIC,
StampMicro or Kitchen for examples. The model is to demonstrate what the reference time looks like so that the Format and Parse methods
can apply the same transformation to a general time value.

Some valid layouts are invalid time values for time.Parse, due to formats such as _ for space padding and Z for zone information.

Within the format string, an underscore _ represents a space that may be replaced by a digit if the following number (a day) has two digits;
for compatibility with fixed-width Unix time formats.

A decimal point followed by one or more zeros represents a fractional second, printed to the given number of decimal places. A decimal point
followed by one or more nines represents a fractional second, printed to the given number of decimal places, with trailing zeros removed.
When parsing (only), the input may contain a fractional second field immediately after the seconds field, even if the layout does not signify
its presence. In that case a decimal point followed by a maximal series of digits is parsed as a fractional second.

Numeric time zone offsets format as follows:

-0700  ±hhmm
-07:00 ±hh:mm
-07    ±hh
Replacing the sign in the format with a Z triggers the ISO 8601 behavior of printing Z instead of an offset for the UTC zone. Thus:

Z0700  Z or ±hhmm
Z07:00 Z or ±hh:mm
Z07    Z or ±hh
The recognized day of week formats are "Mon" and "Monday". The recognized month formats are "Jan" and "January".

The formats 2, _2, and 02 are unpadded, space-padded, and zero-padded day of month. The formats __2 and 002 are space-padded and zero-padded
three-character day of year; there is no unpadded day of year format.

Text in the format string that is not recognized as part of the reference time is echoed verbatim during Format and expected to appear verbatim
in the input to Parse.
*/

// ToGo takes in a datetime format string specified using the POSIX standard
// and converts it to a string in Go's datetime format.
func ToGo(formatString string) (string, error) {
	var out strings.Builder
	var lastPercent bool
	for _, v := range formatString {
		if lastPercent {
			switch v {
			case 'U', 'u', 'V', 'W', 'w', 'X', 'x':
				return "", fmt.Errorf("%%%s not supported in Go: %s", string(v), formatString)
			case '%':
				out.WriteRune('%')
			case 'A':
				out.WriteString("Monday")
			case 'a':
				out.WriteString("Mon")
			case 'B':
				out.WriteString("January")
			case 'b':
				out.WriteString("Jan")
			case 'C':
				out.WriteString("06")
			case 'c':
				out.WriteString("Mon Jan _2 15:04:05 2006")
			case 'D':
				out.WriteString("1/2/06")
			case 'd':
				out.WriteString("2")
			case 'e':
				out.WriteString("_2")
			case 'F':
				out.WriteString("2006-01-02")
			case 'H':
				out.WriteString("15")
			case 'h':
				out.WriteString("Jan")
			case 'I':
				out.WriteString("3")
			case 'j':
				out.WriteString("002")
			case 'k':
				out.WriteString("_15")
			case 'l':
				out.WriteString("_3")
			case 'M':
				out.WriteString("04")
			case 'm':
				out.WriteString("1")
			case 'n':
				out.WriteRune('\n')
			case 'p':
				out.WriteString("PM")
			case 'R':
				out.WriteString("15:04")
			case 'r':
				out.WriteString("3:04:05 PM")
			case 'S':
				out.WriteString("05")
			case 'T':
				out.WriteString("15:04:05")
			case 't':
				out.WriteRune('\t')
			case 'v':
				out.WriteString("_2-Jan-2006")
			case 'Y':
				out.WriteString("2006")
			case 'y':
				out.WriteString("06")
			case 'Z':
				out.WriteString("MST")
			case 'z':
				out.WriteString("-0700")
			default:
				return "", errors.New("invalid format string: " + formatString)
			}
			lastPercent = false
		} else {
			if v == '%' {
				lastPercent = true
			} else {
				out.WriteRune(v)
			}
		}
	}
	return out.String(), nil
}
