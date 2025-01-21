package controllers

import (
	"os"
	"testing"
)

func TestReadValidateInputFile(t *testing.T) {

	// helper function to create a temp file
	createTemp := func(t *testing.T, content string) string {
		t.Helper()

		tempFile, err := os.CreateTemp("", "testfile_*.txt")
		if err != nil {
			t.Fatal(err)
		}

		defer tempFile.Close()

		_, err = tempFile.WriteString(content)
		if err != nil {
			t.Fatal(err)
		}

		return tempFile.Name()
	}
	t.Run("valid input file", func(t *testing.T) {
		content := "3\n2 5 0\n##start\n0 1 2\n##end\n1 9 2\n3 5 4\n0-2\n0-3\n2-1\n3-1\n2-3"
		tempFile := createTemp(t, content)
		defer os.Remove(tempFile)
		expected := []string{
			"3", "2 5 0", "##start", "0 1 2", "##end", "1 9 2", "3 5 4", "0-2", "0-3", "2-1", "3-1", "2-3",
		}
		result, err := ReadValidateInputFile(tempFile)
		if err != nil {
			t.Fatal(err)
		}
		if len(result) != len(expected) {
			t.Fatalf("expected %d lines results, got %d", len(expected), len(result))
		}

		for i := range result {
			if result[i] != expected[i] {
				t.Fatalf("expected %q, got %q", expected[i], result[i])
			}
		}
	})

	// more tests need to be added here.
}

func TestContainsASCII(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  bool
	}{
		{"valid1", "Hello World", true},
		{"invalid1", "こんにちは、世界！", false},
		{"valid2", "be the best", true},
		{"empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsASCII(tt.input); got != tt.want {
				t.Errorf("ContainsASCII() found an invalid character = %v, want %v", got, tt.want)
			}
		})
	}
}
