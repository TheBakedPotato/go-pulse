package pulse

/*
#cgo LDFLAGS: -lpulse-simple
#include <stdlib.h>
#include <pulse/simple.h>
*/
import "C"
import "unsafe"

func Connect() int {
	var pulseSimple *C.pa_simple
	var sampleSpec C.struct_pa_sample_spec
	sampleSpec.format = C.PA_SAMPLE_S16NE
	sampleSpec.channels = 2
	sampleSpec.rate = 44100

	appName := C.CString("Fooapp")
	defer C.free(unsafe.Pointer(appName))
	descStr := C.CString("Music")
	defer C.pa_simple_free(pulseSimple)

	pulseSimple = C.pa_simple_new(
		nil,
		appName,
		C.PA_STREAM_PLAYBACK,
		nil,
		descStr,
		&sampleSpec,
		nil,
		nil,
		nil)
	defer C.free(unsafe.Pointer(pulseSimple))

	if pulseSimple == nil {
		return 0
	}

	return 1
}
