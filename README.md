# tis620
Golang package to handle Thai Industrial Standard 620 (tis620) characters.

## Example
    import "github.com/tupkung/tis620"

    utf_8 := tis620.ToUTF8(tis620bytes)

    tis_620 := tis620.ToTIS620(utf8runes)

## References
* https://th.wikipedia.org/wiki/TIS-620
* https://www.utf8-chartable.de/unicode-utf8-table.pl?start=3584&number=128&utf8=0x

## Credits
* https://github.com/varokas/tis620
* https://github.com/pallat/tis620
