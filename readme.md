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

and the Result type, based on `either.Type`

```golang
r := result.New(123) // result.Type[int]
r.IsError() // false

if w, err := r.Unwrap(); err != nil {
  print(w) // 123
}

result.Map(r, func(t int) string {
  return "123"
}) // result.Type[string]

r.MapError(func(err error) error {
  return fmt.Errorf("faile: %w", err)
}) // equals r, as MapError will only execute when r is error result
```

# License

MIT