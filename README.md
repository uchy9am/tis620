# tis620
Golang package to handle Thai Industrial Standard 620 (tis620) characters.

## Example
    import (
        "fmt"
        "github.com/uchy9am/tis620"
    )

    utf_8 := tis620.ToUTF8(tis620bytes)

    tis_620 := tis620.ToTIS620(utf8runes)

    b, _ := os.ReadFile("utf-8-text.txt")
    chk := tis620.CheckTIS620(b)
    if !chk {
        fmt.Println("no tis-620 file encoding")
    }

## References
* https://th.wikipedia.org/wiki/TIS-620
* https://www.utf8-chartable.de/unicode-utf8-table.pl?start=3584&number=128&utf8=0x

## Credits
* https://github.com/tupkung/tis620 (folk)
* https://github.com/varokas/tis620
* https://github.com/pallat/tis620
