package main

import (
	"testing"
	"os"
)


func BenchmarkArgsWithJoin(b *testing.B) {
	os.Args = []string{"cmd", "arg1", "arg2", "arg3", "arg4", "arg5"}
	for i := 0; i < b.N; i++ {
		argsWithJoin()
    }
}


func BenchmarkArgsWithoutJoin(b *testing.B) {
	os.Args = []string{"cmd", "arg1", "arg2", "arg3", "arg4", "arg5"}
	for i := 0; i < b.N; i++ {
		argsWithoutJoin()
    }
}
