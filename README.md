# pathmatch

a Go generic module that matches proper target(URL(s) etc..) for the basepath.

## Installation

```
    go get github.com/lifthus/pathmatch
```

## Usage

```go
    ...
    PathIntMap := map[string]int{
		"/":      1, // root path / to 1
		"a/a":    2, // /a/a to 2
		"/a/b/":  3, // /a/b to 3
		"/a/b/c": 4, // /a/b/c to 4
		"/a///c": 5, // /a/c to 5
	}
	matcher, err := NewPathMatcher[int](PathIntMap)
	if err != nil {
		t.Errorf("NewPathMatcher has failed: %v", err)
	}

	target, path, ok := matcher.Match("a") // 1, "/a", true
    // root("/") target will be returned, since no basepath is matched.
    // if root path isn't defined, "ok" will be false.
    target, path, ok = matcher.Match("/a/c/def") // 5, "/def", true
    ...
```

Because it is a generic module, you can use any type of target.
