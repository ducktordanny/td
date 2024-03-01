package config

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func getEditor() string {
	if envEditor := os.Getenv("EDITOR"); envEditor != "" {
		return envEditor
	}
	return "vim"
}

func addDefaultContent(content string, file *os.File) {
	if content == "" {
		return
	}
	err := os.WriteFile(file.Name(), []byte(content), 0644)
	if err != nil {
		log.Fatalln("Could not write content to temporary file for edit:", err)
	}
}

func openEditor(file *os.File) {
	editor := getEditor()
	cmd := exec.Command("sh", "-c", editor+" "+file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalln("Couldn't open editor:", err)
	}
}

func getEditResult(file *os.File) string {
	bytes, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatalln("Could not read new content:", err)
	}
	return strings.Trim(string(bytes), "\n")
}

func GetContentByEditor(content string) string {
	file, err := os.CreateTemp("", "td-editor")
	if err != nil {
		log.Fatalln("Couldn't initialize temporary file for editing: ", err)
	}
	defer os.Remove(file.Name())

	addDefaultContent(content, file)
	openEditor(file)
	return getEditResult(file)
}
