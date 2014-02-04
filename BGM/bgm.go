package BGM

/*
#cgo LDFLAGS: -lbass
#include "bass.h"
*/
import "C"
import "unsafe"

var hs C.HSTREAM

func BGM(filename string) {
	C.BASS_Init(-1, 44100, 0, nil, nil)
	hs = C.BASS_StreamCreateFile(C.FALSE, unsafe.Pointer(C.CString(filename)), 0, 0, C.BASS_SAMPLE_LOOP)
	C.BASS_ChannelPlay(C.DWORD(hs), C.TRUE)
}
func Stop() {
	C.BASS_ChannelSetAttribute(C.DWORD(hs), C.BASS_ATTRIB_VOL, 1.0)
	C.BASS_ChannelStop(C.DWORD(hs))
	C.BASS_Free()
}
func Vol(v float32) {
	C.BASS_ChannelSetAttribute(C.DWORD(hs), C.BASS_ATTRIB_VOL, C.float(v))
}
