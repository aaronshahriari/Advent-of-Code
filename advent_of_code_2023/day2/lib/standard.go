package lib

import (
    "os"
    "log"
)

func ReadTxt(file_path string) []byte {
    // read txt in to content (string) variable
    content, err := os.ReadFile(file_path)
    if err != nil {
        log.Fatal(err)
    }

    return content
}

