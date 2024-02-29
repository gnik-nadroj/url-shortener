package encoding

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Base62Encode(n uint64) string {
	encodedValue := ""

    if n == 0 {
        return "0"
    }

    for n > 0 {
        encodedValue = string(base62Chars[n%62]) + encodedValue
        n = n / 62
    }


	return encodedValue;
}