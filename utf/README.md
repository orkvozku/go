<h1 align="center">utf package for golang</h1>

utf is a Golang package for handling UTF-encoded data. It supports UTF-8, UTF-16BE, UTF-16LE, UTF-32BE, UTF-32LE, MUTF-8 (Modified UTF-8), and some other variations. Optionally, a Byte Order Mark (BOM) can be read to detect the encoding.

This module may be considered useful in cases where support for other encodings than just UTF-8 is desired.

## Example

If you are new to the Go language, consider perusing the documentation at [golang.org](https://golang.org/doc/) before proceeding.

Download and install it:

```sh
go get github.com/orkvozku/go/utf

```
Decoding Example:

```go
package main

// import this package:
import "github.com/orkvozku/go/utf"

import (
    "fmt"
)

var data = []byte{0xfe,0xff,0x00,0x48,0x00,0x69,0x00,0x20,0xd8,0x3d,0xde,0x42}

func main() {
    // detect encoding:
    encoding, offset := utf.ParseBom(data)
    if offset > 0 {
        fmt.Printf("Detected a Byte Order Mark (BOM) of %d bytes.\n", offset)
    }
    if encoding == utf.Utf16Be {
        fmt.Println( "The encoding is UTF-16BE." )
    }

    // scan the Unicode code points:
    for {
        r, len := utf.NextRune(data, offset, encoding)
        // alternate: r, len := encoding.NextRune(data, offset)
        if len > 0 {
            fmt.Printf(" U+%06X\n", r)
            offset += len
            continue
        }
        break
    }

    // output:
    // Detected a Byte Order Mark (BOM) of 2 bytes.
    // The encoding is UTF-16BE.
    // U+000048
    // U+000069
    // U+000020
    // U+01F642
}
```
Decoding Example: io.ByteReader to io.RuneReader:
```go
package main

import "github.com/orkvozku/go/utf"
import "io"
import "bytes"
import "fmt"

var data = []byte{0x00,0x48,0x00,0x69,0x00,0x20,0xd8,0x3d,0xde,0x42}

func main() {
    var byteReader = bytes.NewBuffer(data) // provides io.ByteReader interface
    rr := utf.NewReader(byteReader, utf.Utf16Be)
    for {
        r, _, err := rr.ReadRune()
        if err == io.EOF {
            break
        }
        fmt.Printf(" U+%06X\n", r)
    }

    // output:
    //  U+000048
    //  U+000069
    //  U+000020
    //  U+01F642
}
```
Encoding Example:
```go
package main

// import this package:
import "github.com/orkvozku/go/utf"

import (
    "fmt"
)

var text = "Hi \U0001F642"

func main() {
    const enc = utf.Utf8
    for _, r := range text {
        // alternate: replace utf.Bytes(r, enc) with enc.Bytes(r) below
        fmt.Printf("Character \"%c\" encodes to 0x%x\n", r, utf.Bytes(r, enc))
    }

    // output (replace ":)" with however U+01F642 displays on your system):
    // Character "H" encodes to 0x48
    // Character "i" encodes to 0x69
    // Character " " encodes to 0x20
    // Character ":)" encodes to 0xf09f9982
}
```
## Supported Encodings
These encodings are described in multiple specifications including RFC 2781, RFC 3629, RFC 5198, ISO/IEC 10464, and others. Consult the relevant specifications for more details.

* UTF-8:
    - `Utf8` - standard UTF-8 encoding. Overlong encodings, UTF-16 surrogates, and encodings above U+10FFFF are not supported.
    - `Utf8Lax` - similar to `Utf8` except overlong encodings and UTF-16 surrogates are supported. UTF-16 surrogates are combined when they form a matching pair.
    - `Utf8Wild` - similar to `Utf8Lax` except encodings above U+10FFFF all the way to U+1FFFFF are supported.
    - `Mutf8` - similar to `Utf8` except U+000000 is encoded as 2 bytes: 0xC0,0x80. This encoding is specified in [The Java Virtual Machine Specification, Java SE 16 Edition, section 4.4.7](https://docs.oracle.com/javase/specs/jvms/se16/html/jvms-4.html#jvms-4.4.7).
    - `Mutf8Sur` - similar to `Mutf8` except paired UTF-16 surrogates are supported and combined. Characters above U+00FFFF must be encoded via UTF-16 surrogates.
* UTF-16BE:
    - `Utf16Be` - standard UTF-16BE encoding. Un-paired UTF-16 surrogates, and encodings above U+10FFFF are not supported. Paired UTF-16 surrogates are combined.
    - `Utf16BeLax` - similar to `Utf16Be` except un-paired UTF-16 surrogates are supported.
* UTF-16LE:
    - `Utf16Le` - standard UTF-16LE encoding. Un-paired UTF-16 surrogates, and encodings above U+10FFFF are not supported. Paired UTF-16 surrogates are combined.
    - `Utf16LeLax` - similar to `Utf16Le` except un-paired UTF-16 surrogates are supported.
* UTF-32BE:
    - `Utf32Be` - standard UTF-32BE encoding. UTF-16 surrogates and encodings above U+10FFFF are not supported.
    - `Utf32BeLax` - similar to `Utf32Be` except UTF-16 surrogates are supported. UTF-16 surrogates are combined when they form a matching pair.
    - `Utf32BeWild` - similar to `Utf32BeLax` except encodings above U+10FFFF all the way to U+FFFFFFFF are supported.
* UTF-32LE:
    - `Utf32Le` - standard UTF-32LE encoding. UTF-16 surrogates and encodings above U+10FFFF are not supported.
    - `Utf32LeLax` - similar to `Utf16Le` except UTF-16 surrogates are supported. UTF-16 surrogates are combined when they form a matching pair.
    - `Utf32LeWild` - similar to `Utf32LeLax` except encodings above U+10FFFF all the way to U+FFFFFFFF are supported.

The functions, `ParseBom` and `ParseBomLen`, return one of `Utf8`, `Utf16Be`, `Utf16Le`, `Utf32Be`, and `Utf32Le`. To use one of one of the others, you can take the returned value and convert to the encoding of your choice. The helper functions, `ToLax`, `ToWild`, `ToMutf8`, and `ToMutf8Sur` can be used to simplify this task.
```go
    encoding, ofs := utf.ParseBom(data)
    // utf.Utf8 => utf.Utf8Lax, utf.Utf16Le => utf.Utf16LeLax, etc.
    encoding = utf.ToLax(encoding)
```
### Which Encoding to Use
The following encodings are standard and returned by `ParseBom` and `ParseBomLen`. Therefore, it is recommended that you use them unless you have a reason not to. Of these, UTF-8 is the most widely used due to its typically compact representation.

* `Utf8` - UTF-8
* `Utf16Be` - UTF-16BE
* `Utf16Le` - UTF-16LE
* `Utf32Be` - UTF-32BE
* `Utf32Le` - UTF-32LE

Another reason to recommend them is that for each of them, no two byte sequences decode to the same code point. This obstructs some vulnerabilities associated with using alternate encodings of the same text. If the data being decoded is under control of your project (i.e. not provided by external sources), then it is recommended to stick with these or `Mutf8`.

Many real world implementations can produce invalid UTF encodings, such as un-paired UTF-16 surrogates, mixing UTF-16 surrogates into UTF-8 or UTF-32 encodings, and overlong UTF-8 encodings. The "Lax" encodings, `Utf8Lax`, `Utf16BeLax`, `Utf16LeLax`, `Utf32BeLax`, and `Utf32LeLax`, support these. If you want to support UTF encoded data from these implementations, then the "Lax" encodings may be the best choice. `Utf8Lax` also supports MUTF-8. Be aware that these encodings support different byte sequences being decoded to the same code point, which may have security ramifications on your project.

MUTF-8 originated with the Java programming language to provide a variation of UTF-8 that avoids encoding any code point using byte 0x00. This is useful for NUL terminated strings used in a variety of programming languages. The `Utf8Lax` encoding supports these, but two additional encodings are provided that are more strict to MUTF-8: `Mutf8` and `Mutf8Sur`. These encodings change the encoding of U+000000 from the UTF-8 0x00 to the overlong 2-byte encoding 0xC0,0x80. `Mutf8Sur` also changes the encoding of code points in the range, [U+010000,U+10FFFF], to use UTF-16 surrogate pairs.

The "Wild" encodings also support these as well as code points outside the range [U+000000,U+10FFFF]. It is recommended to avoid using these encodings unless these variations must be supported. In most real world applications, the "Lax" encodings are preferrable. The difference is that the "Lax" encodings replace these out-of-range code points with `ReplacementCode` (U+00FFFD), which is typically the appropriate way of handling these.

The following examples consider the code point sequence, U+000048,U+000069,U+000000,U+01F642. When an invalid byte sequence is encountered, the code point, `ReplacementCode` (U+00FFFD), is substituted.

* UTF-8:
    - Correct UTF-8:
        + `0x48,0x69,0x00,0xF0,0x9F,0x99,0x82` - VALID: `Utf8`, `Utf8Lax`, `Utf8Wild`; INVALID: `Mutf8`, `Mutf8Sur`
    - Overlong: `0x48` => `0xC1,0x88`; `0x69` => `0xE0,0x81,0xA9`:
        + `0xC1,0x88,0xE0,0x81,0xA9,0x00,0xF0,0x9F,0x99,0x82` - VALID: `Utf8Lax`, `Utf8Wild`; INVALID: `Utf8`, `Mutf8`, `Mutf8Sur`
    - UTF-16 surrogate pair: U+01F642 -> U+00D83D,U+00DE42 => `0xED,0xA0,0xBD`;`0xED,0xB9,0x82`:
        + `0x48,0x69,0x00,0xED,0xA0,0xBD,0xED,0xB9,0x82` - VALID: `Utf8Lax`, `Utf8Wild`; INVALID: `Utf8`, `Mutf8`, `Mutf8Sur`
    - Un-paired UTF-16 surrogates: U+00DE42,U+00D83D => `0xED,0xB9,0x82`;`0xED,0xA0,0xBD` (VALID encodings decode to the code point sequence, U+000048,U+000069,U+000000,U+00DE42,U+00D83D):
        + `0x48,0x69,0x00,0xED,0xB9,0x82,0xED,0xA0,0xBD` - VALID: `Utf8Lax`, `Utf8Wild`; INVALID: `Utf8`, `Mutf8`, `Mutf8Sur`
    - Out of range (> U+10FFFF): U+110000 => `0xF4,0x90,0x80,0x80` (VALID encodings decdode to the code point sequence, U+000048,U+000069,U+000000,U+110000):
        + `0x48,0x69,0x00,0xF4,0x90,0x80,0x80` - VALID: `Utf8Wild`; INVALID: `Utf8`, `Utf8Lax`, `Mutf8`, `Mutf8Sur`
    - MUTF-8: `0x00` => `0xC0,0x80`:
        + `0x48,0x69,0xC0,0x80,0xF0,0x9F,0x99,0x82` - VALID: `Utf8Lax`, `Utf8Wild`, `Mutf8`; INVALID: `Utf8`, `Mutf8Sur`
    - MUTF-8 with UTF-16 surrogate pair: `0x00` => `0xC0,0x80`; U+01F642 -> U+00D83D,U+00DE42 => `0xED,0xA0,0xBD`;`0xED,0xB9,0x82`:
        + `0x48,0x69,0xC0,0x80,0xED,0xA0,0xBD,0xED,0xB9,0x82` - VALID: `Utf8Lax`, `Utf8Wild`, `Mutf8Sur`; INVALID: `Utf8`, `Mutf8`
    - MUTF-8 with other overlong or unpaired UTF-16 surrogates: `0x00` => `0xC0,0x80`; `0x48` => `0xC1,0x82`; U+00DE42,U+00D83D => `0xED,0xB9,0x82`;`0xED,0xA0,0xBD` (VALID encodings decode to the code point sequence, U+000048,U+000069,U+000000,U+00DE42,U+00D83D):
        + `0xC1,0x88,0xE0,0x81,0xA9,0xC0,0x80,0xED,0xB9,0x82,0xED,0xA0,0xBD` - VALID: `Utf8Lax`, `Utf8Wild`; INVALID: `Utf8`, `Mutf8`, `Mutf8Sur`
* UTF-16BE:
    - Correct UTF-16BE:
        + `0x00,0x48,0x00,0x69,0x00,0x00,0xD8,0x3D,0xDE,0x42` - VALID: `Utf16Be`, `Utf16BeLax`
    - Un-paired UTF-16 surrogates: U+00DE42,U+00D83D (VALID encodings decode to the code point sequence, U+000048,U+000069,U+000000,U+00DE42,U+00D83D):
        + `0x00,0x48,0x00,0x69,0x00,0x00,0xDE,0x42,0xD8,0x3D` - VALID: `Utf16BeLax`; INVALID: `Utf16Be`
* UTF-32LE:
    - Correct UTF-32LE:
        + `0x48,0x00,0x00,0x00,0x69,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x42,0xF6,0x01,0x00` - VALID: `Utf32Le`, `Utf32LeLax`, `Utf32LeWild`
    - UTF-16 surrogate pair: U+01F642 -> U+00D83D,U+00DE42 => `0x3D,0xD8,0x00,0x00`;`0x42,0xDE,0x00,0x00`:
        + `0x48,0x00,0x00,0x00,0x69,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x3D,0xD8,0x00,0x00,0x42,0xDE,0x00,0x00` - VALID: `Utf32LeLax`, `Utf32LeWild`; INVALID: `Utf32Le`
    - Un-paired UTF-16 surrogates: U+00DE42,U+00D83D => `0x42,0xDE,0x00,0x00`;`0x3D,0xD8,0x00,0x00` (VALID encodings decode to the code point sequence, U+000048,U+000069,U+000000,U+00DE42,U+00D83D):
        + `0x48,0x00,0x00,0x00,0x69,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x42,0xDE,0x00,0x00,0x3D,0xD8,0x00,0x00` - VALID: `Utf32LeLax`, `Utf32LeWild`; INVALID: `Utf32Le`
    - Out of range (> U+10FFFF): U+FFFFFFFF => `0xFF,0xFF,0xFF,0xFF` (VALID encodings decdode to the code point sequence, U+000048,U+000069,U+000000,U+FFFFFFFF):
        + `0x48,0x00,0x00,0x00,0x69,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xFF,0xFF,0xFF,0xFF` - VALID: `Utf32LeWild`; INVALID: `Utf32Le`, `Utf32LeLax`

## BOM API
### `utf.ParseBom(data []byte) (utf.Encoding, int)`
Same as ParseBomLen, except that len(data) is not provided.

```go
var data = []byte{0xef, 0xbb, 0xbf, 0x48, 0x69, 0x20, 0xf0, 0x9f, 0x99, 0x82}
encoding, offset := utf.ParseBom(data)  // returns utf.Utf8, 3
encoding = utf.ToLax(encoding) // optional: choose "Lax" encoding
// alternate: encoding = encoding.ToLax()
for {
    r, len := utf.NextRune(data, offset, encoding)
    // alternate: r, len := encoding.NextRune(data, offset)
    if len > 0 {
        fmt.Printf(" U+%06X\n", r)
        offset += len
        continue
    }
    break
}
```

### `utf.ParseBomLen(data []byte, dataLen int) (utf.Encoding, int)`
Parses the Byte Order Mark (BOM) at the start of the data and returns the corresponding standard encoding (i.e. `Utf8`, `Utf16Be`, `Utf16Le`, `Utf32Be`, or `Utf32Le`) and byte length of the BOM. If no BOM was detected, then the returned byte length is 0 and an assumed encoding is returned. Typically, this function is called just before decoding the data with `NextRune` or `NextRuneLen`.

```go
var data = []byte{0xef, 0xbb, 0xbf, 0x48, 0x69, 0x20, 0xf0, 0x9f, 0x99, 0x82}
const dataLen = len(data)
encoding, offset := utf.ParseBomLen(data, dataLen)  // returns utf.Utf8, 3
encoding = encoding.ToLax() // optional: choose "Lax" encoding
// alternate: encoding = utf.ToLax(encoding)
for {
    r, len := utf.NextRuneLen(data, offset, dataLen, encoding)
    // alternate: r, len := encoding.NextRuneLen(data, offset, dataLen)
    if len > 0 {
        fmt.Printf(" U+%06X\n", r)
        offset += len
        continue
    }
    break
}
```

### `utf.ParseBomLenRules1(data []byte, dataLen int) (utf.Encoding, int)`
Similar to ParseBomLen, but with an adjustment to the encoding selection decision process that makes it compliant with Version 1.2.2 of the YASM Specification and any other specification that has equivalent decision requirements for encoding.

The specific difference from ParseBomLen deals with patterns of the form, 0x00.0x?? but are not any of 0x00.0x00.0xFE.0xFF, 0x00.0x00.0x00, or 0x00.0x00.0x00.0x00. For these patterns, ParseBomLen selects UTF-16BE unless the pattern would translate to invalid code points under UTF-16BE but a valid code point according to UTF-32BE, in which case UTF-32BE is selected. ParseBomLenRules1 will select UTF-16BE regardless.

As of this writing, the only assigned Unicode code points for which `ParseBomLenRules` produce a different result than `ParseBomLen` are rarely used code points that are unlikely to be used in typical applications.

### `utf.ParseBomLenEncoding(data []byte, dataLen, encoding utf.Encoding) int`
#### `(utf.Encoding) ParseBomLenEncoding(data []byte, dataLen) int`
Similar to ParseBomLen, except that the encoding is specified. If a BOM corresponding to the specified encoding is detected, the length is returned. Otherwise, a length of 0 is returned.

```go
var data = []byte{0xef, 0xbb, 0xbf, 0x48, 0x69, 0x20, 0xf0, 0x9f, 0x99, 0x82}
const dataLen = len(data)
const encoding = utf.Utf8Lax
offset := encoding.ParseBomLenEncoding(data, dataLen)  // returns 3
for {
    r, len := encoding.NextRuneLen(data, offset, dataLen)
    if len > 0 {
        fmt.Printf(" U+%06X\n", r)
        offset += len
        continue
    }
    break
}
```

### `utf.Bom(encoding utf.Encoding) []byte`
#### `(utf.Encoding) utf.Bom() []byte`
Returns a Byte Order Mark (BOM) for the specified encoding. A splice of zero bytes is returned for unrecognized encodings.

```go
bom := utf.Bom(utf.Utf16Be) // returns []byte{0xfe, 0xff}
```

## Decoding API
### `utf.NextRune(data []byte, offset, encoding utf.Encoding) (rune, int)`
#### `(utf.Encoding) NextRune(data []byte, offset) (rune, int)`
Same as NextRunLen, except that len(data) is not provided.

```go
const data = []byte{0x48, 0x69, 0x20, 0xf0, 0x9f, 0x99, 0x82}
for i := 0; ; i++ {
    r, delta := utf.NextRune(data, i, utf.Utf8)
    if delta == 0 {
        break
    }
    fmt.Printf("Next rune: %v\n", r)
    i += delta
}
```

### `utf.NextRuneLen(data []byte, offset, dataLength, encoding utf.Encoding) (rune, int)`
#### `(utf.Encoding) utf.NextRuneLen(data []byte, offset, dataLength) (rune, int)`
Decodes the bytes starting at the specified offset and encoding into a code point, returning the code point and the number of bytes used to decode the code point. Invalid data returns `ReplacementCode` (U+00FFFD) and the corresponding length of invalid data to skip over. If there is insufficient data to decode another code point, then `ReplacementCode` (U+00FFFD) and a byte length of 0 is returned.

```go
const data = []byte{0x48, 0x69, 0x20, 0xf0, 0x9f, 0x99, 0x82}
const dataLen = len(data)
for i := 0; ; i++ {
    r, delta := utf.NextRuneLen(data, i, dataLen, utf.Utf8)
    if delta == 0 {
        break
    }
    fmt.Printf("Next rune: %v\n", r)
    i += delta
}
```


### `utf.NewReader(br io.ByteReader, encoding utf.Encoding) io.RuneReader`
#### `(utf.Encoding) utf.NewReader(br io.ByteReader) io.RuneReader`
Returns an object with an io.RuneReader interface that decodes the bytes through `utf.NextRuneLen`. If there is a BOM at the start of the data, it will be parsed. If the encoding is set to utf.UtfNotUsed, then `utf.ParseBomLen` is called to parse the encoding; otherwise, `utf.ParseBomLenEncoding` is called to skip over any existing BOM.

```go
// required imports: "bytes", "io"
const data = []byte{0x48, 0x69, 0x20, 0xf0, 0x9f, 0x99, 0x82}
byteReader := bytes.NewBuffer(data) // requires import "bytes"
runeReader := utf.NewReader(byteReader, utf.Utf8)
for {
    r, _, err := runeReader.ReadRune()
    if err == io.EOF { // requires import "io"
        break
    }
    fmt.Printf("Next rune: %v\n", r)
}
```

## Encoding API
### `utf.Bytes(r rune, encoding utf.Encoding) []byte`
#### `(utf.Encoding) utf.Bytes(r rune) []byte`
Encodes a code point (rune) into a slice of bytes of the specified encoding. The "Lax" encodings produce the same results as the standard encodings. Out of range code points are encoded as `ReplacementCode` (U+00FFFD) except if a "Wild" encoding is used that supports code points in that range (i.e. `Utf8Wild` in range [U+110000, U+1FFFFF], `Utf32BeWild`/`Utf32LeWild` in range [U+110000, U+FFFFFFFF]). `Mutf8` produces the same results as `Utf8` except for code point U+000000. `Mutf8Sur` produces the same results as `Mutf8` except that UTF-16 surrogates are used in the range [U+010000, U+10FFFF].

```go
data := utf.Bytes(0x1f642, utf.Utf8)  // returns []byte{0xf0, 0x9f, 0x99, 0x82}
```

## Miscellaneous Utilities Interface
### `utf.ToStandard(encoding utf.Encoding) utf.Encoding`
#### `(utf.Encoding) utf.ToStandard() utf.Encoding`
Converts any encoding to its corresponding base encoding (`Utf8`, `Utf16Be`, `Utf16Le`, `Utf32Be`, or `Utf32Le`).

```go
encoding := utf.ToStandard(utf.Mutf8Sur)  // utf.Mutf8Sur => utf.Utf8
```

### `utf.ToLax(encoding utf.Encoding) utf.Encoding`
#### `(utf.Encoding) utf.ToLax() utf.Encoding`
Converts any encoding to its corresponding "Lax" encoding (`Utf8Lax`, `Utf16BeLax`, `Utf16LeLax`, `Utf32BeLax`, or `Utf32LeLax`).

```go
encoding, ofs := utf.ParseBom(data)
// utf.Utf8 => utf.Utf8Lax, utf.Utf16Le => utf.Utf16LeLax, etc.
encoding = utf.ToLax(encoding)
```

### `utf.ToWild(encoding utf.Encoding) utf.Encoding`
#### `(utf.Encoding) utf.ToWild() utf.Encoding`
Converts any encoding to its corresponding "Wild" encoding (`Utf8Wild`, `Utf32BeWild`, or `Utf32LeWild`). UTF-16 encodings do not have a "Wild" version so they are converted to the corresponding "Lax" version.

```go
encoding, ofs := utf.ParseBom(data)
// utf.Utf8 => utf.Utf8Wild, utf.Utf16Le => utf.Utf16LeLax, utf.Utf32Be => utf.Utf32BeWild, etc.
encoding = utf.ToWild(encoding)
```

### `utf.ToMutf8(encoding utf.Encoding) utf.Encoding`
#### `(utf.Encoding) utf.ToMutf8() utf.Encoding`
Converts UTF-8 and MUTF-8 encodings to `Mutf8`. Other encodings are returned unchanged.

```go
encoding, ofs := utf.ParseBom(data)
// utf.Utf8 => utf.Mutf8, utf.Mutf8Sur => utf.Mutf8, utf.Utf16Le => utf.Utf16Le, etc.
encoding = utf.ToMutf8(encoding)
```

### `utf.ToMutf8Sur(encoding utf.Encoding) utf.Encoding`
#### `(utf.Encoding) utf.ToMutf8Sur() utf.Encoding`
Converts UTF-8 and MUTF-8 encodings to `Mutf8Sur`. Other encodings are returned unchanged.

```go
encoding, ofs := utf.ParseBom(data)
// utf.Utf8 => utf.Mutf8Sur, utf.Mutf8 => utf.Mutf8Sur, utf.Utf16Le => utf.Utf16Le, etc.
encoding = utf.ToMutf8Sur(encoding)
```
