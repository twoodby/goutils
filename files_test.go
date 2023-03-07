package goutils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileExistsStrict(t *testing.T) {
	testDir := t.TempDir()
	_, err := os.Create(fmt.Sprintf("%s/file.txt", testDir))
	assert.NoError(t, err)
	assert.False(t, FileExistStrict(fmt.Sprintf("%s/file_nope.txt", testDir)))
	assert.True(t, FileExistStrict(fmt.Sprintf("%s/file.txt", testDir)))
}

func TestFileExists(t *testing.T) {
	testDir := t.TempDir()
	_, err := os.Create(fmt.Sprintf("%s/file.txt", testDir))
	assert.NoError(t, err)

	e, err := FileExists(fmt.Sprintf("%s/file_nope.txt", testDir))
	assert.False(t, e)
	assert.NoError(t, err)

	e, err = FileExists(fmt.Sprintf("%s/file.txt", testDir))
	assert.True(t, e)
	assert.NoError(t, err)

	// TODO: Write a test to  handle when a error  would be returned maybe something with permissions
}
