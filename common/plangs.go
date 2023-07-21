package common

// Supported programming languages
type PLanguage string

const (
	C   PLanguage = "c"
	CPP           = "cpp"
	CC            = "cc" // C/C++
	// later Circle and CPP2
)

type FileExtensionType string

const (
	HEADER     FileExtensionType = "HEADER"     // .h, .hpp, .hxx
	SOURCE                       = "SOURCE"     // .c, .cpp, .cc
	MODULE                       = "MODULE"     // .cxx
	OBJECT                       = "OBJECT"     // .o, .a, .so
	EXECUTABlE                   = "EXECUTABLE" // .exe, .out, <none>
)

type LibType string

const (
	HEADER_LIB LibType = "INTERFACE" // .h, .hpp, .hxx
	OBJECT_LIB         = "OBJECT"    // .o
	STATIC_LIB         = "STATIC"    // .a, .lib
	SHARED_LIB         = "SHARED"    // .so, .dll
)

var MainTemplate = map[PLanguage]string{
	C: `
// Main function of your binary.
int main(int argc, char *argv[]) 
{

	return 0;
}
	`,
	CPP: `
// Main function of your binary.
int main() 
{

}
	`,
	CC: `
// Main function of your binary.
int main() 
{
	
}
	`, // just like in cpp
}
