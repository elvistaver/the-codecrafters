CLI tool that converts numbers between
  bases. It runs from the terminal like this:

     <> convert 1E hex
       ✦ Decimal: 30

     <> convert 10 bin
       ✦ Decimal: 2

     <> convert 255 dec
       ✦ Binary:  11111111
       ✦ Hex:     FF

  Support Operation:
  → Support three input bases: hex, bin, dec.
  → For dec input, output both binary and hex.
  → For hex and bin input, output only decimal.
  → All hex output must be uppercase.
  → The program keeps running until: quit

  Handles these validation:
  → "1G" is not valid hex.
  → "10201" is not valid binary.
  → "abc" is not a valid decimal.
  → Negative numbers must be supported for dec.
  → Empty input must not crash the program.
