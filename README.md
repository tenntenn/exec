# exec

[![PkgGoDev](https://pkg.go.dev/badge/tenntenn/exec)](https://pkg.go.dev/tenntenn/exec)

`exec` is a command executor.

```go
func Vers(module string) ([]string, error) {
	dir, err := ioutil.TempDir("", "vers*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)

	env := &exec.Env{Dir: dir}
	env.Run("go", "mod", "init", "tmpmodule")
	env.Run("go", "get", module)
	pr, pw := io.Pipe()
	go func() {
		env.Stdout = pw
		env.Run("go", "list", "-m", "-versions", "-json", module)
		pw.Close()
	}()

	var vers struct{ Versions []string }
	err = json.NewDecoder(pr).Decode(&vers)
	if err != nil {
		return nil, err
	}
	if err := env.Err(); err != nil {
		return nil, err
	}
	return vers.Versions, nil
}
```
