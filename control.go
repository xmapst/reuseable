//go:build !linux && !windows

package reuseable

import (
    "syscall"
)

func rawControl(rawConn syscall.RawConn) error {
    return nil
}

// See net.ListenConfig and net.Dialer's Control attribute
func control(network string, address string, rawConn syscall.RawConn) error {
    // See syscall.RawConn.Control
    if err := rawControl(rawConn); err != nil {
        return err
    }
    return nil
}
