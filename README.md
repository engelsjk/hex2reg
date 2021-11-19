# hex2n

```hex2n``` is a port of the JS function [```n_reg```](https://github.com/wiedehopf/tar1090/blob/master/html/registrations.js#L202) that is used in the ADS-B web interface [```wiedehopf/tar1090```](https://github.com/wiedehopf/tar1090). It converts an [ICAO 24-bit Mode S code](https://en.wikipedia.org/wiki/Aviation_transponder_interrogation_modes#ICAO_24-bit_address) (as a base 16 hexadecimal integer) to an N-Number. It is only valid for a United States registered Mode S code.

```go
fmt.Println(Hex2N(0xAA0DB8))
// "N747NA"
```
