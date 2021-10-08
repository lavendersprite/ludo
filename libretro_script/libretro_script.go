package libretro_script

// #cgo CFLAGS: -I${SRCDIR}/../libretro -I${SRCDIR}/libretro_script/include
// #cgo LDFLAGS: -L${SRCDIR}/libretro_script -lretro_script -lm
/*

#include <libretro_script.h>
#include <stdlib.h>

*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/libretro/ludo/state"
)

// this is apparently the universal function pointer type used by cgo.
type fptr = *[0]byte

func free(cs *C.char) {
	C.free(unsafe.Pointer(cs))
}

func Init() {
	// allow libretro_script to intercept core functions
	state.Core.SymRetroSetEnvironment = unsafe.Pointer(
		C.retro_script_intercept_retro_set_environment(fptr(state.Core.SymRetroSetEnvironment)),
	)
	state.Core.SymRetroGetMemorySize = unsafe.Pointer(
		C.retro_script_intercept_retro_get_memory_size(fptr(state.Core.SymRetroGetMemorySize)),
	)
	state.Core.SymRetroGetMemoryData = unsafe.Pointer(
		C.retro_script_intercept_retro_get_memory_data(fptr(state.Core.SymRetroGetMemoryData)),
	)
	state.Core.SymRetroRun = unsafe.Pointer(
		C.retro_script_intercept_retro_run(fptr(state.Core.SymRetroRun)),
	)
}

func LoadScript(scriptPath string) {
	cs := C.CString(scriptPath)
	defer free(cs)
	handle := int(C.retro_script_load(cs))
	if handle == 0 {
		fmt.Printf("retro_script error occurred: ", C.GoString(C.retro_script_get_error()), "\n")
	}
}
 
 
 
