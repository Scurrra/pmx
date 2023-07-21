package dirs

import (
	"fmt"
	"os"

	. "github.com/Scurrra/pmx/common"
)

type Test struct {
	// TODO:
}

// Create `test` directory
func (t Test) CreateDir() error {
	err_dir := os.Mkdir("test", os.ModePerm)
	if err_dir != nil {
		return err_dir
	}

	return nil
}

// Create `test/main.<type>`
func (t Test) CreateTemplate(plang PLanguage) error {
	if _, err := os.Stat("test"); os.IsNotExist(err) {
		err_dir := t.CreateDir()
		if err_dir != nil {
			return err_dir
		}
	}

	f, err_f := os.Create(fmt.Sprintf("test/main.%s", plang))
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
