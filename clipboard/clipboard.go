package clipboard

import (
    "os/exec"

    "github.com/pkg/errors"
)

// Clipboard interface for xclip and xsel utilities
type ClipBoard interface {
    // Write to clipboard
    Copy(text []byte)   error
    // Read from clipboard to output
    Paste()             ([]byte, error)
}

// Generic clipboard struct
type clipboard struct {
}

// NewClipboard returns a new Clipboard interface following
// the given tool
func NewClipBoard(tool string) (ClipBoard, error) {
    var clip ClipBoard

    // Check if tool is installed on system
    if _, err := exec.LookPath("/usr/bin/" + tool); err != nil {
        return nil, err
    }

    if tool == "xclip" {
        clip = newXClipBoard()
        return clip, nil
    }

    //if tool == "xsel" {
        //return newXSelBoard()
    //}

    return nil, errors.Errorf("%s is not supported", tool)
}
