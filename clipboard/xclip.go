package clipboard

import (
    "os/exec"

    //"github.com/pkg/errors"
)

type xclipboard struct {
    CopyCmd             []string
    PasteCmd            []string
}

func newXClipBoard() *xclipboard {
    xclip := &xclipboard{
        CopyCmd: []string{"xclip", "-in", "-selection", "clipboard"},
        PasteCmd: []string{"xclip", "-out", "-selection", "clipboard"},
    }

    return xclip
}

// Copy copies text to the clipboard
func (c *xclipboard) Copy(text []byte) error {
    cmd := exec.Command(c.CopyCmd[0], c.CopyCmd[1:]...)
    in, err := cmd.StdinPipe()
    if err != nil {
        return err
    }

    if err := cmd.Start(); err != nil {
        return err
    }
    if _, err := in.Write([]byte(text)); err != nil {
        return err
    }
    if err := in.Close(); err != nil {
        return err
    }

    return cmd.Wait()
}

// Paste reads from the clipboard and returns text
func (c *xclipboard) Paste() ([]byte, error) {
    cmd := exec.Command(c.PasteCmd[0], c.PasteCmd[1:]...)
    return cmd.Output()
}
