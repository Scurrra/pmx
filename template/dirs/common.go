package dirs

import (
	. "github.com/Scurrra/pmx/common"
)

type CodeFile struct {
	FileHash     [32]byte
	PLang        PLanguage
	Type         FileExtensionType
	Path         string
	Dependencies []string
}
type Headers struct {
	SubDirName string
	Files      []CodeFile
}

type Sources struct {
	SubDirName string
	Files      []CodeFile
}

type Modules struct {
	SubDirName string
	Files      []CodeFile
}

type Binaries struct {
	SubDirName string
	Files      []CodeFile
}
