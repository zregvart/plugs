package main

// #cgo LDFLAGS: -ldl
// #include <stdlib.h>
// #include <dlfcn.h>
//
// int invoke(void *fp, int a, int b)
// {
//   return ((int (*)(int a, int b))fp)(a, b);
// }
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	for _, op := range []string{"add", "subtract"} {
		lib := C.CString(fmt.Sprintf("lib/lib%s.so", op))
		defer C.free(unsafe.Pointer(lib))

		h := C.dlopen(lib, C.RTLD_LAZY)
		if h == nil {
			err := C.dlerror()
			fmt.Printf("%s\n", C.GoString(err))
			continue
		}
		defer C.dlclose(h)

		sym := C.CString(op)
		defer C.free(unsafe.Pointer(sym))

		fp := C.dlsym(h, sym)
		if fp == nil {
			err := C.dlerror()
			fmt.Printf("%s\n", C.GoString(err))
			continue
		}

		fmt.Printf("%s = %d\n", op, C.invoke(fp, 3, 2))
	}
}
