# Contributing to ratelimit-go

## How Can I Contribute?

### Reporting Bugs
- Check the [Issues](https://github.com/dmx3377/ratelimit-go/issues) page to see if the bug has already been reported.
- If not, open a new issue. Clearly describe the problem, include steps to reproduce it, and note your Go version.

### Suggesting Enhancements
- Open an issue to discuss the change before you start coding.
<!-- - to add later when v1.1 released -->

### Pull Requests
1. **Fork** the repository.
2. **Branch**: Create a new branch for your fix or feature (`git checkout -b feature/a-cool-feature`).
3. **Code**: Write your code following standard Go conventions (`gofmt`).
4. **Test**: Ensure all tests pass (`go test ./...`).
5. **Commit**: Keep your commit messages clear and concise.
6. **Push**: Push to your branch and open a Pull Request.

## Coding Standards
- All public functions and types must have comments for `pkg.go.dev` compatibility.
- Use `sync.RWMutex` for thread-safe operations where appropriate.
- Avoid external dependencies unless absolutely necessary.

## License
By contributing, you agree that your contributions will be licensed under the **Apache License 2.0**.
