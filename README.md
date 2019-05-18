# Hasgo

Hasgo is a code generator influenced by Haskell. 
It comes with some functions/types out-of-the-box so you can start using it without running the generator. 

We want to focus on being:
* Immutable 
* Typed (no `interface{}`)
* Nil-safe

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

| Function | Signature          | String | Number | Struct | Maps | Description |
|----------| -----------------  | :----: | :----: | :----: | :--: | ----------- |
| `Sum`    | `[a] -> a`         |   ✓    |   ✓    |        |      | The sum of elements in the slice|

## Contributing

TODO: Write contribution guidelines
