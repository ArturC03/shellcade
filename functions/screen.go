package screen 

import "fmt"
import "fmt"
import "os"
import "syscall"
import "unsafe"

func Clear()  {
    fmt.Print("\033[?25h")
}

func Size() (width, height int, err error) {
  var dimensions [4]
  _, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(os.Stdout.Fd()), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions)), 0, 0, 0)
    if errno != 0 {
        return -1, -1, errno
    }
    return int(dimensions[1]), int(dimensions[0]), nil
}
