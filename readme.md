# Recuerda siempre hacer BET

- Benchmark
- Example
- Test

```
BenchmarFunctionOrMethod(b *testing.B)
ExampleFunctionOrMethod()
TestFunctionOrMethod(t *testing.T)
```

# Comandos

```
godoc -http=localhost:puerto
go test -bench . || -bench ./...
go test -cover
go test -coverprofile c.out ./...
go tool cover -html=c.out
```
