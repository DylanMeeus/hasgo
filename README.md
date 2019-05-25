# Hasgo

Hasgo is a code generator influenced by Haskell. 
It comes with some functions/types out-of-the-box so you can start using it without running the generator. 

We want to focus on being:
* Immutable 
* Typed (no `interface{}`)
* Nil-safe

## Pie

The inspiration for Hasgo, as well as some ideas around implementation come from the lovely [Pie
library](https://github.com/elliotchance/pie), made by [Elliot Chance](https://github.com/elliotchance). 
It's safe to say that Hasgo would not exist without Pie. However, the way Pie and Hasgo work is not
the same and neither is the focus of the project. If you don't find a function in Hasgo, check out
Pie! :smiley:

## Example

```go
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

| Function | Signature                   | String | Number | Struct | Description |
|----------| --------------------------  | :----: | :----: | :----: | ----------- |
| `Abs`    | `[a] -> [a]`                |        |   ✓    |        | Return a slice containing the absolute values|
| `Filter` | `[a] -> (a -> bool) -> [a]` |   ✓    |   ✓    |    ✓   | Filter the slice based on a predicate|
| `Head`   | `[a] -> a`                  |   ✓    |   ✓    |    ✓   | Return the first element|
| `Init`   | `[a] -> [a]`                |   ✓    |   ✓    |    ✓   | Returns all elements minus the last|
| `Last`   | `[a] -> a`                  |   ✓    |   ✓    |    ✓   | Returns the last element|
| `Map`    | `[a] -> (a -> a) -> [a]`    |   ✓    |   ✓    |    ✓   | Returns a slice with the function applied to each element of the input|
| `Maximum`| `[a] -> a`                  |        |   ✓    |        | Returns the largest element|
| `Minimum`| `[a] -> a`                  |        |   ✓    |        | Returns the lowest element|
| `Reverse`| `[a] -> [a]`                |   ✓    |   ✓    |    ✓   | Returns a slice with the elements reversed|
| `sort`   | `[a] -> [a]`                |   ✓    |   ✓    |        | Returns a sorted slice (original remains unsorted)|
| `Sum`    | `[a] -> a`                  |   ✓    |   ✓    |    ✓   | The sum of elements in the slice|
| `Tail`   | `[a] -> [a]`                |   ✓    |   ✓    |    ✓   | Returns all elements minus the first|
| `Uncons` | `[a] -> (a, [a])`           |   ✓    |   ✓    |    ✓   | Returns a tuple of the head and tail of the slice|


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

