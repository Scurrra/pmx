package dirs

import (
	"fmt"
	"os"

	. "github.com/Scurrra/pmx/common"
	. "github.com/eminarican/safetypes"
)

// src directory contents
type SourceCode struct {
	// All private headers
	Headers Option[Headers]

	// Source code files to corresponding headers
	Sources Option[Sources]

	// C++20 module files
	Modules Option[Modules]

	// All types of allowed files in src/
	CodeFiles []CodeFile
}

// Create empty directory `src` with subdirectories
func (sc SourceCode) CreateDirs() error {
	err_dir := os.Mkdir("src", os.ModePerm)
	if err_dir != nil {
		return err_dir
	}

	err_dir = os.Mkdir("headers", os.ModePerm)
	if err_dir != nil {
		return err_dir
	}

	err_dir = os.Mkdir("sources", os.ModePerm)
	if err_dir != nil {
		return err_dir
	}

	err_dir = os.Mkdir("modules", os.ModePerm)
	if err_dir != nil {
		return err_dir
	}

	return nil
}

// Create empty directory `src` with subdirectories and the `main` file
func (sc SourceCode) CreateTemplateBin(plang PLanguage) error {
	if _, err := os.Stat("src"); os.IsNotExist(err) {
		err_dir := sc.CreateDirs()
		if err_dir != nil {
			return err_dir
		}
	}

	f, err_f := os.Create(fmt.Sprintf("src/main.%s", plang))
	if err_f != nil {
		return err_f
	}
	defer f.Close()

	_, err_f = f.Write([]byte(MainTemplate[plang]))
	if err_f != nil {
		return err_f
	}

	return nil
}
