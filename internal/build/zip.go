package build

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// CreateZip creates a zip archive containing all files and folders except:
// - .git
// - .gitignore
// - greengrass-build/
// - recipe.yaml
// - gdk-config.json
func CreateZip(dir string, zipPath string) {
	log.Printf("Creating zip file: %s\n", zipPath)

	// inpiered from: https://stackoverflow.com/a/63233911

	file, err := os.Create("output.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Crawling: %#v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			// add a trailing slash for creating dir
			path = fmt.Sprintf("%s%c", path, os.PathSeparator)
			_, err = w.Create(path)
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}

		if info.Name() == ".git" || info.Name() == ".gitignore" || info.Name() == "greengrass-build" || info.Name() == "recipe.yaml" || info.Name() == "gdk-config.json" {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		f, err := w.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}

	err = filepath.Walk(dir, walker)
	if err != nil {
		log.Fatal(err)
	}

}
