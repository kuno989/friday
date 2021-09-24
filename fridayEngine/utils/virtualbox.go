package utils

import "os/exec"

func MemoryDump(vmname, path string) error {
	filename := path + "/memory.cap"
	cmd := exec.Command("VBoxManage", "debugvm", vmname, "dumpvmcore", "--filename="+filename)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
