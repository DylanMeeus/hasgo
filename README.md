# Hasgo [![Build Status](https://travis-ci.com/DylanMeeus/hasgo.svg?branch=master)](https://travis-ci.com/DylanMeeus/hasgo)

Coverage status: [gocover.io](https://gocover.io/github.com/DylanMeeus/hasgo/types)
Our report card[Report Card](https://goreportcard.com/badge/github.com/DylanMeeus/hasgo)](https://goreportcard.com/report/github.com/DylanMeeus/hasgo)

Hasgo is a code generator with functions influenced by Haskell. 
It comes with some types out-of-the-box so you can start using it without running the generator.
Specifically you can start using Hasgo's `Strings` and `Ints` types. 

We want to focus on being:
* Immutable 
* Strongly-Typed (no `interface{}`)
* Nil-safe

## Pie

The inspiration for Hasgo, as well as some ideas around implementation come from the lovely [Pie
library](https://github.com/elliotchance/pie), made by [Elliot Chance](https://github.com/elliotchance). 
It's safe to say that Hasgo would not exist without Pie. However, the way Pie and Hasgo work is not
the same and neither is the focus of the project. If you don't find a function in Hasgo, check out
Pie! :smiley:

## Example

```go
import . "github.com/DylanMeeus/hasgo/types"
func EpicFunction() {
	// create a range of -10 -> 10. Take the absolute values, keep only even numbers, and sum them.
	result := IntRange(-10,10).
		Abs().
		Filter(func(i int64) bool {
			return i % 2 == 0
		}).
		Sum()
	// result = 60 
}
```

You can find more [examples here](https://github.com/DylanMeeus/hasgo/tree/master/examples). 

## Installation

```bash
go get -u github.com/DylanMeeus/hasgo
```

## Types 
* `Ints` ([]int64)
* `Strings` ([]string)

## Functions

These are the function currently available with Hasgo.
It shows you which type of data they operate on as well as the Haskell type definition. 
The first symbol of the signature is actually the method receiver in Go terms. 

Alternatively, you can consult the [godoc](https://godoc.org/github.com/DylanMeeus/hasgo/types)

### Generic functions
These functions can be generated for every type.

| Function     | Signature                   | String | Number | Struct | Description |
|------------  | --------------------------  | :----: | :----: | :----: | ----------- |
| `Abs`        | `[a] -> [a]`                |        |   ✓    |        | Return a slice containing the absolute values|
| `All`        | `[a] -> (a -> bool) -> bool`|   ✓    |   ✓    |    ✓   | Returns true if the predicate applies to all elements in the slice|
| `Any`        | `[a] -> (a -> bool) -> bool`|   ✓    |   ✓    |    ✓   | Returns true if one or more elements satisfy the predicate|
| `Average`    | `[a] -> a`                  |        |   ✓    |        | Returns the average of all elements|
| `Delete `    | `[a] -> a -> [a]`           |   ✓    |  ✓     |    ✓   | Returns the slice with the first occurance of the element deleted.|
| `Filter`     | `[a] -> (a -> bool) -> [a]` |   ✓    |   ✓    |    ✓   | Filter the slice based on a predicate|
| `Head`       | `[a] -> a`                  |   ✓    |   ✓    |    ✓   | Return the first element|
| `Init`       | `[a] -> [a]`                |   ✓    |   ✓    |    ✓   | Returns all elements minus the last|
| `Intercalate`| `[a] -> [[a]] -> [a]`       |   ✓    |   ✓    |    ✓   | Intersperses the slice in between the provided 2d-slice |
| `Intersperse`| `[a] -> a -> [a]`           |   ✓    |   ✓    |    ✓   | Intersperses the value in between all elements of the provided slice|  
| `Last`       | `[a] -> a`                  |   ✓    |   ✓    |    ✓   | Returns the last element|
| `Length`     | `[a] -> int`                |   ✓    |   ✓    |    ✓   | Returns the length of the slice|
| `Map`        | `[a] -> (a -> a) -> [a]`    |   ✓    |   ✓    |    ✓   | Returns a slice with the function applied to each element of the input|
| `Maximum`    | `[a] -> a`                  |        |   ✓    |        | Returns the largest element|
| `Minimum`    | `[a] -> a`                  |        |   ✓    |        | Returns the lowest element|
| `Modes`      | `[a] -> [a]`                |   ✓    |   ✓    |    ✓   | Returns the elements with the highest frequency |
| `Nub`        | `[a] -> [a]`                |   ✓    |   ✓    |    ✓   | Returns a Slice containing one of each of the input elements |
| `Null`       | `[a] -> bool`               |   ✓    |   ✓    |    ✓   | Returns true if the slice is empty, false otherwise|
| `Product`    | `[a] -> a`                  |   ✓    |        |        | Returns the product of all elements in the slice.|
| `Reverse`    | `[a] -> [a]`                |   ✓    |   ✓    |    ✓   | Returns a slice with the elements reversed|
| `Sort`       | `[a] -> [a]`                |   ✓    |   ✓    |        | Returns a sorted slice (original remains unsorted)|
| `Sum`        | `[a] -> a`                  |   ✓    |   ✓    |    ✓   | The sum of elements in the slice|
| `Tail`       | `[a] -> [a]`                |   ✓    |   ✓    |    ✓   | Returns all elements minus the first|
| `Take`       | `[a] -> uint64 -> [a]`      |   ✓    |   ✓    |    ✓   | Take N elements from the slice, or all if N exceeds the length.|
| `Uncons`     | `[a] -> (a, [a])`           |   ✓    |   ✓    |    ✓   | Returns a tuple of the head and tail of the slice|
| `Unlines`    | `[a] -> string`             |   ✓    |   ✓    |    ✓   | Returns a newline separated string of all elements in the slice| 
| `Unwords`    | `[a] -> string`             |   ✓    |   ✓    |    ✓   | Returns a space-separated string of all elements in the slice| 


### Hardcoded functions

The built-in types (Strings, Ints) have some functions defined on them that are not generated.
Mostly because we could not create them in a generic way. 


| Type     | Function         | Signature                    | Description  |
|--------- |---------------   |------------------------------|--------------| 
| `Ints`   | `Equals`         | \*`Ints -> Ints -> bool`       | Returns true if both slices contain the same elements| 
| `Ints`   | `EqualsOrdered`  | \*`Ints -> Ints -> bool`       | Returns true if both slices contain the same elements, in the same position| 
| `Ints`   | `IntRange`       | `int64 -> int64 -> Ints`     | Return an integer range from [start,stop]|
| `Ints`   | `IntReplicate`   | `uint64 -> int64 -> Ints`    | Return a slice with the input element repeated n times|
| `Strings`| `Equals`         | \*`Strings -> Strings -> bool` | Returns true if both slices contain the same elements| 
| `Strings`| `EqualsOrdered`  | \*`Strings -> Strings -> bool` | Returns true if both slices contain the same elements, in the same position| 
| `Strings`| `StringReplicate`| `uint64 -> string -> Strings`| Return a slice with the input element repeated n times|

\* (Functions prefixed by a star are functions added to the type itself, where first element in the
signature is the method receiver. So for examples, the Equals method is `Ints{1,2}.Equals(Ints{1})`.
But, the IntRange function looks like `hasgo.IntRange(0,10)`.


## Contributing

You can help out Hasgo in a variety of ways! 
Here are some ideas:

* Use Hasgo! :smiley:
* Spread the word (Write a blog, tweet, talk about..)
* Suggest features (Create an issue to make a suggestion)
* Report bugs (Similarly, create an issue)
* Contribute code. (Create a PR, we'll gladly take a look and help you get it merged!)
	* We have separate [contribution guidelines](CONTRIBUTING.md)
	
## What's in a name?
The name Hasgo is a portmanteau of "Haskell" and "Go". I'm a big fan of both languages, though they
are quite different. It's impossible to write real Haskell-like code in Go. There are some obvious
differences between the languages in terms of syntax. I hope the functions in this library stay as
close as possible to their Haskell implementations. There might be extra functions in here that are
not in Haskell, and there _will_ be functions in Haskell that you won't find here.

The inspiration mainly shows in the naming of functions. If the functions were named after Java
lambdas, it'd be called "Jago". Sorry if you expected more Haskell goodness (I'm open to suggestions
of how more haskell in Hasgo!)

