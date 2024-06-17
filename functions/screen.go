package cursor 

import "fmt"
import "os"
import "syscall"
import "unsafe"

func Clear() {
    fmt.Print("\033[2J")
}

func MakeOutline(x,y,width,height int) {
    if width < 2 || height < 2 {
        fmt.Println("Width and height should be at least 2 to draw a box.")
        return
    }

    // Draw top border
    Position(x,y)
    fmt.Print("┌")
    for i := 1; i < width - 1; i++ {
        fmt.Print("─")
    }
    fmt.Println("┐")

    // Draw sides
    for i := 1; i < height - 1; i++ {
        Position(x,y + i)
        fmt.Print("│")
        Position(x + width - 1,y + i)
        fmt.Println("│")
    }

    // Draw bottom border
    Position(x, y + height - 1)
    fmt.Print("└")
    for i := 1; i < width - 1; i++ {
        fmt.Print("─")
    }
    fmt.Println("┘")
}

func GetScreenSize() (width, height int, err error) {
    var dimensions [4]uint16
    _, _, errno := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(os.Stdout.Fd()), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions)), 0, 0, 0)
    if errno != 0 {
        return -1, -1, errno
    }
    return int(dimensions[1]), int(dimensions[0]), nil
}
