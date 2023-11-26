# lo-chain
`lo-chain` is a wrapper of [samber/lo](https://github.com/samber/lo) to use chain method.

# Usage
Create chain instance to use `chain.NewXxx()` and it make it possible to use chain method.
if you want to original value, use `Value()` method.

```go
type comment struct {
	user  string
	title string
}

comments := map[string][]comment{
	"post1": {
		comment{"Jon", "No-title"},
	},
	"post2": {},
	"post3": {
		comment{"Ken", "No-title"},
	},
}

// print posts which have a comment.
chain.NewMapChain(comments).
	PickBy(func(post string, comments []comment) bool {
		return len(comments) > 0
	}).
	Keys().
	ForEach(func(post string, _ int) {
		fmt.Println(post)
	})
```
