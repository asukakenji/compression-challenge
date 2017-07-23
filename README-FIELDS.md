# Fields

## First Field

Since the first field is a time within a day, it ranges from `0` (`00:00:00`) to
`86439` (`23:59:59`).

Note that:
- 2 ** 16 = 65536
- 2 ** 17 = 131072

Therefore, the time could be represented in a 17-bit field.

Previously:
- A variable bit-length scheme (like UTF-8) could have been applied to use less
  bits for smaller values and more bits for larger values. However, it seems that
  it is not a good idea, since the time value does not, in general, flavor smaller values.

Now:
- A variable bit-length scheme is available as a built-in package in Golang (`"encoding/binary"`).

### Updated
The maximum difference between this line's timestamp and previous line's timestamp is `39` seconds.

Note that:
- 2 ** 5 = 32
- 2 ** 6 = 64

Therefore, the time diff could be represented in a 6-bit field.

## Second Field

Since it is always the same as the first field with the "seconds" part removed.
There is no need to encode it in the output.

## Third, Fourth, Fifth Field

These fields are of similar values.
Therefore they could be stored using a diff scheme.

This part is not used:
- field3 larger than field4: 6665 times
- field4 larger than field3: 63691 times
- field3 the same as field4: 163 times

The decimal number is converted into an integer by multiplying it by `1000` to
ease handling. For example, `108.009` is treated as `108009`.
It is done by parsing the integral part and the decimal part separately,
instead of using the built-in floating point parser of Golang.
Floating point is not a suitable choice since a lost of precision may occur
when dealing with decimal numbers.

The smallest number is `108.009`, it is encoded as `0`.
The largest number is `109.672`, it is encoded as `1663`.

Note that:
- 2 ** 10 = 1024
- 2 ** 11 = 2048

Therefore, each of these fields could be represented in a 11-bit field.

## Sixth Field

It could be seen as a enum of the following:
1. DUBA
1. FFS
1. FXCM
1. FXDC
1. FXDD
1. GAIN
1. KZ
1. PEP
1. PFD
1. SBD
1. SEBS
1. TDFX

Note that:
- 2 ** 3 = 8
- 2 ** 4 = 16

Therefore, the enum could be represented in a 4-bit field.
