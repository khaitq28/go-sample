package main

/*
Files Tutorial
-------------
This tutorial covers file operations in Go:
1. Basic File Operations
   - Reading files
   - Writing files
   - File permissions
2. Advanced Features
   - Buffered I/O
   - File seeking
   - Temporary files
3. Best Practices
   - Error handling
   - Resource cleanup
   - File locking
*/

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// FileManager handles file operations
type FileManager struct {
	baseDir string
}

// NewFileManager creates a new file manager
func NewFileManager(baseDir string) (*FileManager, error) {
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, fmt.Errorf("error creating base directory: %v", err)
	}
	return &FileManager{baseDir: baseDir}, nil
}

// WriteFile writes content to a file
func (fm *FileManager) WriteFile(filename string, content []byte) error {
	filepath := filepath.Join(fm.baseDir, filename)
	return ioutil.WriteFile(filepath, content, 0644)
}

// ReadFile reads content from a file
func (fm *FileManager) ReadFile(filename string) ([]byte, error) {
	filepath := filepath.Join(fm.baseDir, filename)
	return ioutil.ReadFile(filepath)
}

// AppendFile appends content to a file
func (fm *FileManager) AppendFile(filename string, content []byte) error {
	filepath := filepath.Join(fm.baseDir, filename)
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	if _, err := f.Write(content); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}

// CopyFile copies a file from source to destination
func (fm *FileManager) CopyFile(src, dst string) error {
	srcPath := filepath.Join(fm.baseDir, src)
	dstPath := filepath.Join(fm.baseDir, dst)

	sourceFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("error creating destination file: %v", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("error copying file: %v", err)
	}
	return nil
}

// ListFiles lists all files in a directory
func (fm *FileManager) ListFiles(dir string) ([]string, error) {
	dirPath := filepath.Join(fm.baseDir, dir)
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %v", err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames, nil
}

// CreateTempFile creates a temporary file
func (fm *FileManager) CreateTempFile(prefix string) (*os.File, error) {
	return ioutil.TempFile(fm.baseDir, prefix)
}

func main() {
	// Create file manager
	fm, err := NewFileManager("test_files")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll("test_files") // Clean up after example

	// Write a file
	content := []byte("Hello, World!")
	if err := fm.WriteFile("hello.txt", content); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created hello.txt")

	// Read the file
	readContent, err := fm.ReadFile("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read content: %s\n", string(readContent))

	// Append to the file
	appendContent := []byte("\nThis is appended content.")
	if err := fm.AppendFile("hello.txt", appendContent); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Appended to hello.txt")

	// Copy the file
	if err := fm.CopyFile("hello.txt", "hello_copy.txt"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Copied hello.txt to hello_copy.txt")

	// List files
	files, err := fm.ListFiles(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files in directory:")
	for _, file := range files {
		fmt.Printf("- %s\n", file)
	}

	// Create a temporary file
	tempFile, err := fm.CreateTempFile("temp_")
	if err != nil {
		log.Fatal(err)
	}
	defer tempFile.Close()

	fmt.Printf("Created temporary file: %s\n", tempFile.Name())

	// Demonstrate buffered reading
	if err := fm.WriteFile("numbers.txt", []byte("1\n2\n3\n4\n5")); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(filepath.Join(fm.baseDir, "numbers.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Reading numbers.txt line by line:")
	for scanner.Scan() {
		fmt.Printf("Line: %s\n", scanner.Text())
	}

	// Demonstrate file seeking
	file.Seek(0, 0)
	reader := bufio.NewReader(file)
	fmt.Println("Reading first character:")
	firstChar, err := reader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("First character: %c\n", firstChar)
}
