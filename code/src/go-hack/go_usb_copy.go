package main

import (
	"io"
	"strings"
	"os"
	"path/filepath"
	"fmt"
	"strconv"
	"syscall"
	"golang.org/x/sys/windows/registry"
)

// 获取系统中所有盘符
func GetSystemDisks() []string {
	 // 获取系统dll
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	 // 获取dll中函数
	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")

	n, _, _ := GetLogicalDrives.Call()

	s := strconv.FormatInt(int64(n), 2)
	var allDrives = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:",
		 "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：",
		 "U：", "V：", "W：", "X：", "Y：", "Z："}
	temp := allDrives[0:len(s)]	 
	var d []string
    for i, v := range s {
        if v == 49 {
            l := len(s) - i - 1
            d = append(d, temp[l])
        }
    }
    var drives []string
    for i, v := range d {
        drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
    }
    return drives
}

// 获取插入的U盘盘符
func GetUDisk() []string {
    //查询注册表，判断是否插入U盘
    k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\USBSTOR\Enum`, registry.QUERY_VALUE)
    if err != nil {
        fmt.Println("Not have U-Disk")
        return nil
    }
    defer k.Close()
    // 获取注册表中值，得到插入了几个U盘
    count, _, err := k.GetIntegerValue("Count")
    // 获取全部盘符
    disks := GetSystemDisks()

    return disks[len(disks)-int(count):]
}

//递归复制目录
func copyDir(src string, dest string)  {
	src_original := src
	err := filepath.Walk(src, func(src string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			//			fmt.Println(f.Name())
			//copyDir(f.Name(), dest+"/"+f.Name())
		} else {
			//fmt.Println(src)
			//fmt.Println(src_original)
			//fmt.Println(dest)
 
			dest_new := strings.Replace(src, src_original, dest, -1)
			//fmt.Println(dest_new)
			//fmt.Println("CopyFile:" + src + " to " + dest_new)
			CopyFile(src, dest_new)
		}
		//println(path)
		return nil
	})
	if err != nil {
		//fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}


func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()

	//fmt.Println("dst:" + dst)
	dst_slices := strings.Split(dst, "\\")
	dst_slices_len := len(dst_slices)
	dest_dir := ""
	for i := 0; i < dst_slices_len-1; i++ {
		dest_dir = dest_dir + dst_slices[i] + "\\"
	}
	//dest_dir := getParentDirectory(dst)
	//fmt.Println("dest_dir:" + dest_dir)
	b, err := PathExists(dest_dir)
	if b == false {
		err := os.Mkdir(dest_dir, os.ModePerm) //在当前目录下生成md目录
		if err != nil {
			fmt.Println(err)
		}
	}
	dstFile, err := os.Create(dst)
 
	if err != nil {
		fmt.Println(err.Error())
		return
	}
 
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		//println(path)
		return nil
	})
	if err != nil {
		//fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


func main() {
	if len(GetUDisk()) <= 0 {
		fmt.Println("没有插入U盘")
		return
	}

	for _, v := range  GetUDisk() {
		fmt.Println(v + "\\")
		copyDir(v + "\\","F:\\go_copy")
	}
}