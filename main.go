package hash_colour

import (
   "fmt"
   "image/color"
   "math/bits"

   "github.com/cespare/xxhash"
   gc "github.com/lucasb-eyer/go-colorful"
)

const MASK5BIT = 0b11111
const MASK6BIT = 0b111111
const MASK7BIT = 0b1111111
const MASK8BIT = 0xFF
const MASK24BIT = 0xFFFFFF
const MASK32BIT = 0xFFFFFFFF

const SAT_FLOOR = 28
const LUV_FLOOR = 28

func Colour5532(s string) (sat uint64, luv uint64, hue float64) {
   hash := xxhash.Sum64String(s)
   sat = take_bits(&hash, MASK5BIT)
   luv = take_bits(&hash, MASK5BIT)
   hue = float64(take_bits(&hash, MASK32BIT) & MASK32BIT) / float64(MASK32BIT) * 360
   return
}

func Colour5524(s string) (sat uint64, luv uint64, hue float64) {
   hash := xxhash.Sum64String(s)
   sat = take_bits(&hash, MASK5BIT)
   luv = take_bits(&hash, MASK5BIT)
   hue = float64(take_bits(&hash, MASK24BIT) & MASK24BIT) / float64(MASK24BIT) * 360
   return
}

func HSLuv(s string) (c color.Color) {
   sat, luv, hue := Colour5532(s)
   c = gc.HSLuv(hue, float64(sat+SAT_FLOOR)/100, float64(luv+LUV_FLOOR)/100)
   return
}

func HSLuv24(s string) (c color.Color) {
   sat, luv, hue := Colour5524(s)
   c = gc.HPLuv(hue, float64(sat+LUV_FLOOR)/100, float64(luv+LUV_FLOOR)/100)
   return
}

func HexColour(c color.Color) (hex string) {
   r, g, b, _ := c.RGBA()
   hex = "#" + fmt.Sprintf("%0x", r>>8) + fmt.Sprintf("%0x", g>>8) + fmt.Sprintf("%0x", b>>8)
   return
}

func take_bits(hash *uint64, mask uint64) (result uint64) {
   result = *hash & mask
   h := bits.RotateLeft64(*hash, -bits.OnesCount64(mask))
   hash = &h
   return
}

