# Manage Static Assests with `embed` (Golang 1.16)

## Overview

Golang 1.16 new package `embed` manages static assets that will be embedded in the application binary. Any files from a package or package subdirectory can be "embedded" and retrived as variable of type `string` or `bytes[]`.

```
import _ "embed"

//go:embed hello.txt
var s string

//go:embed hello.txt
var b []bytes
```

In addition, you can retrive you embeded file with a variable of type `FS`. This let you define a pattern that macht which file need to be "embeded" in your application 

```
import "embed"

//go:embed assets/*
var f embed.FS
```

[Official Documentation](https://golang.org/pkg/embed/)

## Use embeded to 


## 3 Articles a related subject

* [Working with Embed in Go 1.16 Versions](https://lakefs.io/working-with-embed-in-go/)
* [Golang embed static assets in binary (with React build example)](https://www.akmittal.dev/posts/go-embed-files/)
* [How to Use //go:embed](https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/)