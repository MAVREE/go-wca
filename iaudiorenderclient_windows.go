// +build windows

package wca

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func arcGetBuffer(arc *IAudioRenderClient, requiredBufferSize uint32, data interface{}) (err error) {
	dataValue := reflect.ValueOf(data).Elem()
	hr, _, _ := syscall.Syscall(
		arc.VTable().GetBuffer,
		3,
		uintptr(unsafe.Pointer(arc)),
		uintptr(requiredBufferSize),
		dataValue.Addr().Pointer())
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}

func arcReleaseBuffer(arc *IAudioRenderClient, writtenBufferSize, flag uint32) (err error) {
	hr, _, _ := syscall.Syscall(
		arc.VTable().ReleaseBuffer,
		3,
		uintptr(unsafe.Pointer(arc)),
		uintptr(writtenBufferSize),
		uintptr(flag))
	if hr != 0 {
		err = ole.NewError(hr)
	}
	return
}
