# dsu

[![CI]](https://github.com/RealA10N/dsu/actions/workflows/ci.yaml)
[![codecov]](https://codecov.io/gh/RealA10N/dsu)
[![pre-commit.ci status]](https://results.pre-commit.ci/latest/github/RealA10N/dsu/main)

Disjoint Set Union (Union Find) *performance focused* 🔥 implementation in Golang.

```go
func ExampleDsu() {
    d := dsu.NewDsu(3)
    d.Union(0, 1)

    if d.Find(0) == d.Find(1) {
        fmt.Println("0 and 1 are in the same set")
    }

    if d.Find(1) != d.Find(2) {
        fmt.Println("1 and 2 are in different sets")
    }

    if d.Size(0) == 2 {
        fmt.Println("Set with the element 0 has 2 elements")
    }

    // Output:
    // 0 and 1 are in the same set
    // 1 and 2 are in different sets
    // Set with the element 0 has 2 elements
}
```

[ci]: https://github.com/RealA10N/dsu/actions/workflows/ci.yaml/badge.svg
[codecov]: https://codecov.io/gh/RealA10N/dsu/graph/badge.svg
[pre-commit.ci status]: https://results.pre-commit.ci/badge/github/RealA10N/dsu/main.svg
