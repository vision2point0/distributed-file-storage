package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Upload File (API 1)
func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // limit upload size
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// File split logic
	chunks := splitFile(file, 1024*1024) // 1 MB chunks

	// Parallel uploading
	var wg sync.WaitGroup
	fileID := generateFileID()
	for i, chunk := range chunks {
		wg.Add(1)
		go func(i int, chunk []byte) {
			defer wg.Done()
			saveChunkToDB(fileID, i, chunk)
		}(i, chunk)
	}
	wg.Wait()

	w.Write([]byte(fmt.Sprintf("File uploaded successfully with ID: %s", fileID)))
}

// Get File Metadata (API 2)
func getFileMetadataHandler(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Path[len("/files/"):]
	metadata, err := fetchFileMetadata(fileID)
	if err != nil {
		http.Error(w, "Error fetching file metadata", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(metadata)
}

// Download File (API 3)
func downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Path[len("/download/"):]
	chunks, err := fetchFileChunks(fileID)
	if err != nil {
		http.Error(w, "Error fetching file", http.StatusInternalServerError)
		return
	}

	// Merge chunks
	mergedFile := mergeChunks(chunks)

	w.Header().Set("Content-Disposition", "attachment; filename=merged_file")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(mergedFile)
}
