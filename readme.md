# cround

A floating point number rounding package for Go which calls the rounding functions from the C standard library.

The Go standard library does not provide a rounding function in the math package, which requires developers to write their own. Here are some rounding functions you can use instead of doing that.

The default rounding mode is "banker's rounding". That is, values are rounded to the nearest integer value of x, with ties rounded to even integers.

For a pure Go version of this functionality, see the [round](http://github.com/mhale/round) package.

## Usage

Using this package may require a C compiler to be installed on the system, which is not ideal for all developers. However it does allow Go programs to make use of the special case handling in mature C implementations to theoretically improve accuracy.

Import cround as you would any other package, and then call one of the exposed functions. See the source code for a listing.

If you are not sure which function to use, you should use Round. A more explicit function is ToNearestEven.

### Round

Round is a convenience wrapper for ToNearestEven.

```
value := cround.Round(2.5)
// value is 2.0
```

### RoundTo

RoundTo is a convenience wrapper for ToNearestEven, which rounds to a specified number of decimal places.

```
value := cround.Round(1234.5678, 2)
// value is 1234.57
```

### ToNearestEven

ToNearestEven rounds to the nearest integer value, with ties rounded to even integers.

```
value := cround.ToNearestEven(2.5)
// value is 2.0
```

## Miscellaneous

These links may be useful to learn about floating point rounding modes.

* [IEEE Standard for Floating-Point Arithmetic: Rounding rules (IEEE 754)](https://en.wikipedia.org/wiki/IEEE_floating_point#Rounding_rules)
* [The GNU C Library: Rounding](https://www.gnu.org/software/libc/manual/html_node/Rounding.html)
* [RoundingMode (Java Platform SE 8)](https://docs.oracle.com/javase/8/docs/api/java/math/RoundingMode.html)

## Testing

The tests cover a range of test values, including the zero and infinity special cases. NaN is not included because a comparison between two NaN values always returns false. 

Please submit any additional test cases that demonstrate errors or inaccuracies.

## Licensing

This code is in the public domain. Anyone is free to copy, modify, publish, use, compile, sell, or distribute the original cround code, either in source code form or as a compiled binary, for any purpose, commercial or non-commercial, and by any means.