package clipboard

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestXClipBoard(t *testing.T) {
    var text string = "This is just a clipboard test"

    clip, err := NewClipBoard("toto")
    assert.NotNil(t, err)

    clip, err = NewClipBoard("xclip")
    assert.Nil(t, err)

    //err = clip.Copy([]byte(text))
    //assert.Nil(t, err)

    paste, err := clip.Paste()
    assert.Nil(t, err)
    //assert.EqualValues(t, paste, text)

    fmt.Println("Copy text to clipboard: ", string(text))
    fmt.Println("Paste text from clipboard: ", string(paste))
}
