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

## Installation

```bash
go get -u github.com/DylanMeeus/hasgo`
```

## Types 
* `Ints` ([]int64)
* `Strings` ([]string)

## Functions

These are the function currently available with Hasgo.
It shows you which type of data they operate on as well as the Haskell type definition.

| Function | Signature                   | String | Number | Struct | Maps | Description |
|----------| --------------------------  | :----: | :----: | :----: | :--: | ----------- |
| `Filter` | `[a] -> (a -> bool) -> [a]` |   ✓    |   ✓    |        |      | Filter the slice based on a predicate|
| `Head`   | `[a] -> a`                  |   ✓    |   ✓    |        |      | Return the first element|
| `Init`   | `[a] -> [a]`                |   ✓    |   ✓    |        |      | Returns all elements minus the last|
| `Last`   | `[a] -> a`                  |   ✓    |   ✓    |        |      | Returns the last element|
| `Sum`    | `[a] -> a`                  |   ✓    |   ✓    |        |      | The sum of elements in the slice|
| `Tail`   | `[a] -> [a]`                |   ✓    |   ✓    |        |      | Returns all elements minus the first|

## Contributing

You can help out Hasgo in a variety of ways! 
Here are some ideas:

* Use Hasgo. Just using it will fill you with ideas for the next options! :smiley:
* Spread the word (Write a blog, tweet, talk about..)
* Suggest features (Create an issue to make a suggestion)
* Report bugs (Similarly, create an issue)
* Contribute code. (Create a PR, we'll gladly take a look and help you get it merged!)
	* We have separate [contribution guidelines](CONTRIBUTING.md)


TODO: Write contribution guidelines
