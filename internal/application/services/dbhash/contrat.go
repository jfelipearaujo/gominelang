package dbhash

import "time"

type FileHash struct {
	ID        int64      `json:"id"`
	FilePath  string     `json:"file_path"`
	Hash      string     `json:"hash"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Service interface {
	Open() error
	Close() error

	Exists(filePath string) (*FileHash, error)
	Compare(fileHash *FileHash, compareTo string) (bool, error)
	Store(filePath string) error
}
