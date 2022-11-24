package compress

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func DoZip(srcDir string, zipFileName string) {
	os.RemoveAll(zipFileName)
	zipfile, _ := os.Create(zipFileName)
	defer zipfile.Close()
	archive := zip.NewWriter(zipfile)
	defer archive.Close()
	filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {
		if path == srcDir {
			return nil
		}
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, srcDir+`\`)
		if info.IsDir() {
			header.Name += `/`
		} else {
			header.Method = zip.Deflate
		}
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
}
