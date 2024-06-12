package cursor 

import {
  "fmt"
  "os"
  "syscall"
  "unsafe"
}

func Clear() {
    fmt.Print("\033[2J")
}

func MakeOutline(x,y,width,heigth int) {
    if width < 2 || height < 2 {
        fmt.Println("Width and height should be at least 2 to draw a box.")
        return
    }

    // Draw top border
    cursor.Position(x,y)
    fmt.Print("┌")
    for i := 0; i < width-2; i++ {
        fmt.Print("─")
    }
    fmt.Println("┐")

    // Draw sides
    for i := 0; i < height-2; i++ {
        cursor.Position(x,y + 1)
        fmt.Print("│")
        cursor.Position(x + width,y + 1)
        fmt.Println("│")
    }

    // Draw bottom border
    cursor.Position(x, y + height)
    fmt.Print("└")
    for i := 0; i < width-2; i++ {
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
