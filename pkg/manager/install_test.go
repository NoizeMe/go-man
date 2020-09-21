package manager

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"

	"github.com/NoizeMe/go-man/pkg/releases"
	"github.com/NoizeMe/go-man/pkg/tasks"
)

func TestGoManager_Install(t *testing.T) {
	validVersion := version.Must(version.NewVersion("1.15.2"))
	invalidVersion := version.Must(version.NewVersion("42.1337.3"))

	tempDir := t.TempDir()

	sut, err := NewManager(&tasks.Task{
		ErrorExitCode: 1,
		Output:        os.Stdout,
		Error:         os.Stderr,
	}, tempDir)
	assert.NoError(t, err)
	assert.NotNil(t, sut)

	assert.Error(t, sut.Install(invalidVersion, runtime.GOOS, runtime.GOARCH, releases.IncludeAll))

	assert.NoError(t, sut.Install(validVersion, runtime.GOOS, runtime.GOARCH, releases.IncludeAll))
	assert.DirExists(t, filepath.Join(tempDir, fmt.Sprintf("go%s", validVersion)))

	assert.Error(t, sut.Install(validVersion, runtime.GOOS, runtime.GOARCH, releases.IncludeAll))
	assert.DirExists(t, filepath.Join(tempDir, fmt.Sprintf("go%s", validVersion)))
}