package dbhash_test

import (
	"os"
	"testing"
	"time"

	"github.com/jfelipearaujo/gominelang/internal/application/services/dbhash"
	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	t.Run("Should return true if the file exists", func(t *testing.T) {
		// Arrange
		os.Remove("gominelang.db")

		file := "./testdata/testfile.txt"

		sut := dbhash.New()

		err := sut.Open()
		assert.NoError(t, err)
		defer sut.Close()

		err = sut.Store(file)
		assert.NoError(t, err)

		// Act
		res, err := sut.Exists(file)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Should return false if the file does not exists", func(t *testing.T) {
		// Arrange
		os.Remove("gominelang.db")

		file := "./testdata/testfile.txt"

		sut := dbhash.New()

		err := sut.Open()
		assert.NoError(t, err)
		defer sut.Close()

		// Act
		res, err := sut.Exists(file)

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, res)
	})
}

func TestCompare(t *testing.T) {
	t.Run("Should return true if the hashs are the same", func(t *testing.T) {
		// Arrange
		os.Remove("gominelang.db")

		file := "./testdata/testfile.txt"

		sut := dbhash.New()

		err := sut.Open()
		assert.NoError(t, err)
		defer sut.Close()

		err = sut.Store(file)
		assert.NoError(t, err)

		fileHash, err := sut.Exists(file)
		assert.NoError(t, err)

		// Act
		res, err := sut.Compare(fileHash, file)

		// Assert
		assert.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("Should return false if the hashs are not the same", func(t *testing.T) {
		// Arrange
		os.Remove("gominelang.db")

		file := "./testdata/testfile.txt"

		sut := dbhash.New()

		err := sut.Open()
		assert.NoError(t, err)
		defer sut.Close()

		err = sut.Store(file)
		assert.NoError(t, err)

		// Act
		res, err := sut.Compare(&dbhash.FileHash{
			Hash: "1234567890",
		}, file)

		// Assert
		assert.NoError(t, err)
		assert.False(t, res)
	})
}

func TestStore(t *testing.T) {
	t.Run("Should store a file", func(t *testing.T) {
		// Arrange
		os.Remove("gominelang.db")

		file := "./testdata/testfile.txt"

		sut := dbhash.New()

		err := sut.Open()
		assert.NoError(t, err)
		defer sut.Close()

		// Act
		err = sut.Store(file)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should update a file that already exists", func(t *testing.T) {
		// Arrange
		os.Remove("gominelang.db")

		file := "./testdata/testfile.txt"

		newContent := []byte("this is the content of test file")
		err := os.WriteFile(file, newContent, 0644)
		assert.NoError(t, err)

		sut := dbhash.New()

		err = sut.Open()
		assert.NoError(t, err)
		defer sut.Close()

		err = sut.Store(file)
		assert.NoError(t, err)

		time.Sleep(1 * time.Second)

		newContent = []byte("this is a new content of test file")
		err = os.WriteFile(file, newContent, 0644)
		assert.NoError(t, err)

		// Act
		err = sut.Store(file)

		// Assert
		assert.NoError(t, err)
	})
}
