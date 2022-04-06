# posix_time

The Go date time format specifier language is (as far as I can tell) one's favorite format. The POSIX standard provides
a perfectly cromulent way to specify times and dates across multiple languages. This library allows you to write a 
date time template using the much more common POSIX % escapes and translates it to what Go expects.

This is still pretty simple and needs far better code coverage, but supports what I need right now.

## Usage

```go
posixFormat := "%d-%b-%y"
myTime := time.Date(2021, time.June, 20, 0, 0, 0, 0, time.UTC)
goFormat, err := posix_time.ToGo(posixFormat)
if err != nil {
    panic("unexpected")
}
fmt.Println(myTime.Format(goFormat)) // Expected 20-Jun-21
```

## Limitations

Since `posix_time` does a simple mapping from the POSIX date time format to Go's 
idiosyncratic date time format, some POSIX escapes aren't valid. In particular:

- Go's date time format doesn't support week of year (`%U`, `%V`, `%W`) 
- Go's date time format doesn't support numeric days of the week (`%u`, `%w`)

Using them in the formatting string returns an error.

Because Go's `time.Format` supports only English names for months and days, there is no support for locale. 
That means `%c`, `%X`, and `%x` format dates and times using the POSIX locale standards:
- `%c` is equivalent to `%a %b %e %H:%M:%S %Y`
- `%X` is equivalent to `%H:%M:%S`
- `%x` is equivalent to `%m/%d/%y`

If you want to output translated names for months, days of week, and AM/PM, I recommend you use the https://github.com/goodsign/monday library

## License

This library is licensed under the MIT license.

