package system

import (
        "bytes"
        "os/exec"
)

//execute cmd and return cmd's output
func ExeCmd(s string) (string, error) {
        cmd := exec.Command("/bin/bash", "-c", s)

        var out bytes.Buffer
        cmd.Stdout = &out

        err := cmd.Run()
        return out.String(), err
}
