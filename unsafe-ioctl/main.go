package main

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

func main() {
	//ethInfo, err := unix.IoctlEthtoolDrvinfo()
	f, err := os.Open("/dev/vhost-net")
	if err != nil {
		throw(err)
	}
	defer f.Close()
	var cid uint32
	err = ioctl(f.Fd(), unix.IOCTL_VM_SOCKETS_GET_LOCAL_CID, unsafe.Pointer(&cid))
	if err != nil {
		throw(err)
	}
	fmt.Println(cid)
}

func throw(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(0b01)
}

func ioctl(fd uintptr, request int, argp unsafe.Pointer) error {
	_, _, errno := unix.Syscall(
		unix.SYS_IOCTL,
		fd,
		uintptr(request),
		uintptr(argp),
	)
	if errno != 0 {
		return os.NewSyscallError("ioctl", fmt.Errorf("%d", int(errno)))
	}
	return nil
}
