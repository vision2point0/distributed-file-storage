package main

import (
	"io"
	"log"

	"github.com/google/uuid"
)

// Function to split file into chunks
func splitFile(file io.Reader, chunkSize int64) [][]byte {
	var chunks [][]byte
	buffer := make([]byte, chunkSize)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error reading file:", err)
			break
		}
		chunks = append(chunks, buffer[:n])
	}
	return chunks
}

// Function to save file chunk to database
func saveChunkToDB(fileID uuid.UUID, chunkNumber int, chunkData []byte) {
	_, err := db.Exec("INSERT INTO file_chunks (file_id, chunk_number, chunk_data) VALUES ($1, $2, $3)", fileID, chunkNumber, chunkData)
	if err != nil {
		log.Println("Error saving chunk to database:", err)
	}
}

// Function to fetch file metadata
func fetchFileMetadata(fileID string) ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT chunk_number FROM file_chunks WHERE file_id = $1", fileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metadata []map[string]interface{}
	for rows.Next() {
		var chunkNumber int
		err := rows.Scan(&chunkNumber)
		if err != nil {
			return nil, err
		}
		metadata = append(metadata, map[string]interface{}{
			"chunk_number": chunkNumber,
		})
	}
	return metadata, nil
}

// Function to fetch file chunks from database
func fetchFileChunks(fileID string) ([][]byte, error) {
	rows, err := db.Query("SELECT chunk_data FROM file_chunks WHERE file_id = $1 ORDER BY chunk_number", fileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chunks [][]byte
	for rows.Next() {
		var chunkData []byte
		err := rows.Scan(&chunkData)
		if err != nil {
			return nil, err
		}
		chunks = append(chunks, chunkData)
	}
	return chunks, nil
}

// Function to merge file chunks back into one file
func mergeChunks(chunks [][]byte) []byte {
	merged := []byte{}
	for _, chunk := range chunks {
		merged = append(merged, chunk...)
	}
	return merged
}

// Generate unique file ID
func generateFileID() uuid.UUID {
	return uuid.New()
}
