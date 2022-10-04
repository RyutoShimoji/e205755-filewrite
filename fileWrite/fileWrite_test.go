package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("e205755")
    want := "Hi, e205755. Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}