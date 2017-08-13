package repltwo

/*
REPL that uses raw mode...
It might be a better idea to expose terminal/raw mode as a language lib...

 */

import (
	"os"
	"syscall"
	"unsafe"
	"io"
)

// Terminos and structure from: https://github.com/creack/termios
const (
	getTermios = syscall.TIOCGETA
	setTermios = syscall.TIOCSETA
)

type Termios struct {
	Iflag  uint64
	Oflag  uint64
	Cflag  uint64
	Lflag  uint64
	Cc     [20]byte
	Ispeed uint64
	Ospeed uint64
}

func TcSetAttr(fd uintptr, termios *Termios) error {
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(setTermios), uintptr(unsafe.Pointer(termios))); err != 0 {
		return err
	}
	return nil
}

func TcGetAttr(fd uintptr) (*Termios, error) {
	var termios = &Termios{}
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, getTermios, uintptr(unsafe.Pointer(termios))); err != 0 {
		return nil, err
	}
	return termios, nil
}

func CfMakeRaw(termios *Termios) {
	termios.Iflag &^= (syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP | syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON)
	termios.Oflag &^= syscall.OPOST
	termios.Lflag &^= (syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN)
	termios.Cflag &^= (syscall.CSIZE | syscall.PARENB)
	termios.Cflag |= syscall.CS8
	termios.Cc[syscall.VMIN] = 1
	termios.Cc[syscall.VTIME] = 0
}

func MakeRaw(fd uintptr) (*Termios, error) {
	old, err := TcGetAttr(fd)
	if err != nil {
		return nil, err
	}

	new := *old
	CfMakeRaw(&new)

	if err := TcSetAttr(fd, &new); err != nil {
		return nil, err
	}
	return old, nil
}

func Start() {

	bin, err := syscall.Open("/dev/tty", 0x0, 0)
	raw := os.NewFile(uintptr(bin), "raw")
	rout, _ := MakeRaw(uintptr(bin))
	bout, err := os.OpenFile("/dev/tty", 0x1, 0)
	counter := 0
	for {
		message := make([]byte, 1)
		io.ReadFull(raw, message)
		// If special. Might want to use switch instead of if.
		if message[0] == byte(27) {
			message := make([]byte, 2)
			io.ReadFull(raw, message)

			// UP Key as delete.
			if message[0] == '[' && message[1] == 'A' {
				if counter > 0 {
					for i := 0; i < counter; i++ {
						// Create byte buffer for back-space(delete last) back(move cursor back)
						// It's probably a better idea abstract cursor movement.
						deleteKey := make([]byte, 7)
						deleteKey[0] = byte(27)
						deleteKey[1] = byte('[')
						deleteKey[2] = byte('D')
						deleteKey[3] = byte(' ')
						deleteKey[4] = byte(27)
						deleteKey[5] = byte('[')
						deleteKey[6] = byte('D')
						bout.Write(deleteKey)
					}
				}
				counter = 0
			}
			//bout.Write(message)
		} else if message[0] == '\r' {
			// Return exits
			// Setting tc attr back to previous is important since staying in raw mode
			// Can leave to some nasty side effects.
			TcSetAttr(uintptr(bin), rout)
			syscall.Close(bin)
			bout.Close()
			os.Exit(0)
		} else {
			counter++
			bout.Write(message)
		}
		if err != nil {
			print(err)
		}
	}
	bout.Close()
}