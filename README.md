# Golang Challenges

[![Build Status](https://travis-ci.org/geektimus/go-challenges.svg?branch=master)](https://travis-ci.org/geektimus/go-challenges)

This repo will contain the solutions using Golang for challenges found on:

* Books
* Hacker Rank
* Codility
* CodeFights
* Others

## Challenges

### Others (Challenges proposed by Friends or Solved on hackatons)

#### Recursively reverse a string

Given a string, reverse the characters and return.

## Appendix

### Run the test for all the challenges

```bash
go test -v ./...
```

### Run tests with coverage

```bash
go test -cover
```

### Output the coverage result to a file

```bash
go test -coverprofile=coverage.out
```

### View the test results

Console `go tool cover -func=coverage.out`

HTML `go tool cover -html=coverage.out`
