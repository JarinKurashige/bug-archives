package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func promptSingleLine(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func promptMultiline(prompt string, reader *bufio.Reader) string {
	fmt.Println(prompt)
	fmt.Println("(Type your text. When finished, enter a single line with END.)")

	var lines []string
	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimRight(line, "\n")
		if line == "END" {
			break
		}
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

func reviewForm(title, tags, summary, symptoms, rootCause, fix, detect string) {
	fmt.Println("\n===== REVIEW BUG ENTRY =====")
	fmt.Println("1. Title:")
	fmt.Println(title)
	fmt.Println("\n2. Tags:")
	fmt.Println(tags)
	fmt.Println("\n3. Summary:")
	fmt.Println(summary)
	fmt.Println("\n4. Symptoms:")
	fmt.Println(symptoms)
	fmt.Println("\n5. Root Cause:")
	fmt.Println(rootCause)
	fmt.Println("\n6. Fix:")
	fmt.Println(fix)
	fmt.Println("\n7. How To Detect:")
	fmt.Println(detect)
	fmt.Println("\n============================")
}

func main() {
	fmt.Printf("100%% ChatGPT generated script\n\n")

	reader := bufio.NewReader(os.Stdin)

	// Fields stored here so they can be edited
	title := ""
	tagsInput := ""
	summary := ""
	symptoms := ""
	rootCause := ""
	fix := ""
	detect := ""

	// First pass input
	title = promptSingleLine("Bug Title: ", reader)
	tagsInput = promptSingleLine("Tags (comma separated): ", reader)
	summary = promptMultiline("\nSummary:", reader)
	symptoms = promptMultiline("\nSymptoms:", reader)
	rootCause = promptMultiline("\nRoot Cause:", reader)
	fix = promptMultiline("\nFix:", reader)
	detect = promptMultiline("\nHow To Detect This Bug In The Future:", reader)

	// Editing loop
	for {
		// Format tags for display
		var tags []string
		for _, t := range strings.Split(tagsInput, ",") {
			tag := strings.TrimSpace(t)
			if tag != "" {
				tags = append(tags, "#"+strings.ToLower(strings.ReplaceAll(tag, " ", "_")))
			}
		}
		tagString := strings.Join(tags, " ")

		// Show full review
		reviewForm(title, tagString, summary, symptoms, rootCause, fix, detect)

		fmt.Print("\nEdit a field? (1-7, or 's' to save): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			title = promptSingleLine("\nBug Title: ", reader)
		case "2":
			tagsInput = promptSingleLine("\nTags (comma separated): ", reader)
		case "3":
			summary = promptMultiline("\nSummary:", reader)
		case "4":
			symptoms = promptMultiline("\nSymptoms:", reader)
		case "5":
			rootCause = promptMultiline("\nRoot Cause:", reader)
		case "6":
			fix = promptMultiline("\nFix:", reader)
		case "7":
			detect = promptMultiline("\nHow To Detect This Bug In The Future:", reader)
		case "s", "S":
			// Save and exit
			// Build final tags string
			var ts []string
			for _, t := range strings.Split(tagsInput, ",") {
				tag := strings.TrimSpace(t)
				if tag != "" {
					ts = append(ts, "#"+strings.ToLower(strings.ReplaceAll(tag, " ", "_")))
				}
			}
			tagString := strings.Join(ts, " ")

			// Build filename
			safeTitle := strings.ToLower(title)
			safeTitle = strings.ReplaceAll(safeTitle, " ", "-")
			safeTitle = strings.ReplaceAll(safeTitle, "/", "-")

			date := time.Now().Format("2006-01-02")
			filename := fmt.Sprintf("%s-%s.md", date, safeTitle)

			// Create directory
			baseDir := "bug-archive/bugs"
			os.MkdirAll(baseDir, 0755)

			outputPath := filepath.Join(baseDir, filename)

			// Build markdown
			content := fmt.Sprintf(`# %s

## Summary
%s

## Symptoms
%s

## Root Cause
%s

## Fix
%s

## How To Detect This Bug In the Future
%s

## Tags
%s
`, title, summary, symptoms, rootCause, fix, detect, tagString)

			// Write file
			err := os.WriteFile(outputPath, []byte(content), 0644)
			if err != nil {
				fmt.Println("Error writing file:", err)
				return
			}

			fmt.Println("\nSaved:", outputPath)
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

