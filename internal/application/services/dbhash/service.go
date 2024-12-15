package dbhash

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type service struct {
	dbName string

	db *sql.DB
}

func New() Service {
	return &service{
		dbName: "gominelang.db",
	}
}

func (s *service) Open() error {
	createTables := false
	if _, err := os.Stat(s.dbName); os.IsNotExist(err) {
		createTables = true
	}

	var err error

	s.db, err = sql.Open("sqlite3", s.dbName)
	if err != nil {
		return fmt.Errorf("failed to open database '%s': %w", s.dbName, err)
	}

	if createTables {
		createFileHashTable := `
			CREATE TABLE IF NOT EXISTS file_hash (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				file_path TEXT NOT NULL,
				hash TEXT NOT NULL,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME
			);
		`

		if _, err := s.db.Exec(createFileHashTable); err != nil {
			return fmt.Errorf("failed to create table 'file_hash': %w", err)
		}
	}

	return nil
}

func (s *service) Close() error {
	return s.db.Close()
}

func (s *service) Exists(filePath string) (*FileHash, error) {
	var fileHash FileHash
	err := s.db.QueryRow("SELECT * FROM file_hash WHERE file_path = ?", filePath).
		Scan(&fileHash.ID,
			&fileHash.FilePath,
			&fileHash.Hash,
			&fileHash.CreatedAt,
			&fileHash.UpdatedAt,
		)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	return &fileHash, err
}

func (s *service) Compare(fileHash *FileHash, compareTo string) (bool, error) {
	data, err := os.ReadFile(compareTo)
	if err != nil {
		return false, fmt.Errorf("failed to read file '%s': %w", compareTo, err)
	}

	hash := fmt.Sprintf("%x", sha256.Sum256(data))

	return fileHash.Hash == hash, nil
}

func (s *service) Store(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file '%s': %w", filePath, err)
	}

	hash := fmt.Sprintf("%x", sha256.Sum256(data))

	var fileHash FileHash
	err = s.db.QueryRow("SELECT * FROM file_hash WHERE file_path = ?", filePath).
		Scan(&fileHash.ID,
			&fileHash.FilePath,
			&fileHash.Hash,
			&fileHash.CreatedAt,
			&fileHash.UpdatedAt,
		)
	if err != nil && err == sql.ErrNoRows {
		_, err := s.db.Exec("INSERT INTO file_hash (file_path, hash) VALUES (?, ?)", filePath, hash)
		if err != nil {
			return fmt.Errorf("failed to insert file hash '%s': %w", filePath, err)
		}
		return err
	}

	if fileHash.Hash != hash {
		_, err = s.db.Exec("UPDATE file_hash SET hash = ?, updated_at = ? WHERE file_path = ?", hash, time.Now(), filePath)
		if err != nil {
			return fmt.Errorf("failed to update file hash '%s': %w", filePath, err)
		}
	}

	return nil
}
