package hash_colour

import (
   "testing"

   "github.com/stretchr/testify/assert"
)

func Test_HSLuv(t *testing.T) {
   c := HSLuv("Omnikron13")
   r, g, b, a := c.RGBA()
   hx := HexColour(c)


   assert.Equal(t, 0x5538, int(r))
   assert.Equal(t, 35626, int(g))
   assert.Equal(t, 23522, int(b))
   assert.Equal(t, 0xffff, int(a))
   assert.Equal(t, "#558b5b", hx)
}

func Test_Hex(t *testing.T) {
   c := HSLuv("Omnikron13")
   hx := HexColour(c)

   assert.Equal(t, "#558b5b", hx)

   c = HSLuv("dlvhdr")
   hx = HexColour(c)

   assert.Equal(t, "#bd587e", hx)
}

func Test_HSLuv24(t *testing.T) {
   c := HSLuv24("Omnikron13")
   r, g, b, a := c.RGBA()
   hx := HexColour(c)

   assert.Equal(t, uint32(0x6f9e), uint32(r))
   assert.Equal(t, uint32(0x851b), uint32(g))
   assert.Equal(t, uint32(0x680f), uint32(b))
   assert.Equal(t, uint32(0xffff), uint32(a))
   assert.Equal(t, "#6f8568", hx)
}
