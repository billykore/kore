package path

import "os"

func GetAbsolutePath(file string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	if dir != "/" {
		return dir + "/" + file
	}
	return dir + file
}
