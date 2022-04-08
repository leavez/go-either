# go-either

An Either type for golang with generics

```golang
e := either.NewLeft[int, string](1) // either.Type[int, string]

l, r, isLeft := e.Unwrap()
if isLeft {  // true
  print(l)   // l is 1
} else {
  print(r)
}

e.EitherDo(func(l int) {
  print(l)
}, func(r string) {
  print(r)
})

either.MapLeft(e, func(left int) *int {
  return &left
}) // either.Type[*int, string]

// either.MapRight
// either.Map
```

# License

MIT