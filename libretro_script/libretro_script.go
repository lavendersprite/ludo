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
	C.retro_script_init()

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
	state.Core.SymRetroInit = unsafe.Pointer(
		C.retro_script_intercept_retro_init(fptr(state.Core.SymRetroInit)),
	)
	state.Core.SymRetroDeinit = unsafe.Pointer(
		C.retro_script_intercept_retro_deinit(fptr(state.Core.SymRetroDeinit)),
	)
	state.Core.SymRetroRun = unsafe.Pointer(
		C.retro_script_intercept_retro_run(fptr(state.Core.SymRetroRun)),
	)
	state.Core.SymRetroSetInputPoll = unsafe.Pointer(
		C.retro_script_intercept_retro_set_input_poll(fptr(state.Core.SymRetroSetInputPoll)),
	)
	state.Core.SymRetroSetInputState = unsafe.Pointer(
		C.retro_script_intercept_retro_set_input_state(fptr(state.Core.SymRetroSetInputState)),
	)
}

func Deinit() {
	C.retro_script_deinit()
}

func LoadScript(scriptPath string) {
	cs := C.CString(scriptPath)
	defer free(cs)
	handle := int(C.retro_script_load_lua(cs))
	if handle == 0 {
		fmt.Printf("retro_script error occurred: ", C.GoString(C.retro_script_get_error()), "\n")
	}
}
