hash-colour
===========

Generate colours from hashes for repeatable username colouring, etc.

Basic usage:

```go
import hc "github.com/omnikron13/hash-colour"

c := hc.HSLuv("Omnikron13")

hx := hc.HexColour(c)

println(hx)
```
output: `#558b5b`
```

API Docs
--------

 - `HSLuv(name string) (sat, luv uint64, hue float64)`: Generate a colour from a string.
 - `HexColour(c color.Color) (hex string)`: Convert a color.Color to a hex colour string.
 - `c.RGBA() (r, g, b, a float64)`: Convert HSLuv to RGBA.

