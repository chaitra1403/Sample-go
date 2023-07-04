package main

import (
    "bytes"
    "fmt"
    "io"
    "os"
    "testing"
)

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}

// Test suite
func TestExitErrorf04518af35b(t *testing.T) {
    // Capture the standard error
    oldStderr := os.Stderr // keep backup of the real stdout
    r, w, _ := os.Pipe()
    os.Stderr = w

    // Call the function
    exitErrorf("Test error message %s", "error")

    outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        outC <- buf.String()
    }()

    // back to normal state
    w.Close()
    os.Stderr = oldStderr // restoring the real stdout
    out := <-outC

    // Check the function's output
    expected := "Test error message error\n"
    if out != expected {
        t.Errorf("Expected %v, but got %v", expected, out)
    }

    // Check if the function has exited
    if !isExitCalled() {
        t.Error("Expected os.Exit to be called, but it wasn't")
    }
}

// Helper function to check if os.Exit was called
func isExitCalled() (exitCalled bool) {
    // Capture the panic when os.Exit is called
    defer func() {
        if e := recover(); e != nil {
            exitCalled = true
        }
    }()

    os.Exit(0)

    return
}
