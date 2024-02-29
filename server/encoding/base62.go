package internal_encoding

const base62Chars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const minFiveCharValue uint64 = 62*62*62*62

func Base62Encode(n uint64) string {
	encodedValue := ""

    n += minFiveCharValue

    if n == 0 {
        return "0"
    }

    for n > 0 {
        encodedValue = string(base62Chars[n%62]) + encodedValue
        n = n / 62
    }


	return encodedValue;
}