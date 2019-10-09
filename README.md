# 🏗 Env
Enivroment vars simple. Usefull when you use docker containers.

```go
//check key "host" and if it empty use default value
env.String("host", "https://192.168.0.1")
env.Int("port", 9999)
env.Bool("cache", true)
env.Bool("cache", 1)
env.Float("lat", -1.42556)
env.Strings("hosts", "," , "https://192.168.0.1", "https://192.168.0.1")
```

Useless benchmarks)
```
BenchmarkString-8       13758784                93.9 ns/op             0 B/op          0 allocs/op
BenchmarkInt-8           7840081               149 ns/op              48 B/op          1 allocs/op
BenchmarkStrings-8       8104532               149 ns/op              16 B/op          1 allocs/op
BenchmarkBool-8         13488634                96.1 ns/op             0 B/op          0 allocs/op
```
