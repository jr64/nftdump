package out

import (
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
)

const logo = "\n" +
	"                 --------------------- = ---------------------                 \n" +
	"                 -------------------- === --------------------                 \n" +
	"                 ----------------- ========= -----------------                 \n" +
	"                 ------------- ================= -------------                 \n" +
	"                 --------- ========================= ---------                 \n" +
	"                 ----- ============ NFTdump ============ -----                 \n" +
	"                 --------- ========================= ---------                 \n" +
	"                 ------------- ================= -------------                 \n" +
	"                 ----------------- ========= -----------------                 \n" +
	"                 -------------------- === --------------------                 \n" +
	"                 --------------------- = ---------------------                 \n" +
	"\n"

func Logo() {
	fmt.Print(logo)
}
func InfoPrefix() string {
	return color.FgBlue.Render("[*] ")
}

func ErrorPrefix() string {
	return color.FgRed.Render("[-] ")
}

func SuccessPrefix() string {
	return color.FgGreen.Render("[+] ")
}

func print(prefix string, format string, a ...interface{}) (n int, err error) {
	str := fmt.Sprintf("%s%s\n", prefix, format)
	return fmt.Printf(str, a...)
}

func Info(format string, a ...interface{}) (n int, err error) {
	return print(InfoPrefix(), format, a...)
}

func Success(format string, a ...interface{}) (n int, err error) {
	return print(SuccessPrefix(), format, a...)
}

func Warn(format string, a ...interface{}) (n int, err error) {
	return print(ErrorPrefix(), format, a...)
}

func Fatal(format string, a ...interface{}) {
	print(ErrorPrefix(), format, a...)
	os.Exit(1)
}

func Hierarchical(indent int, format string, a ...interface{}) (n int, err error) {
	str := fmt.Sprintf("| %s- %s\n", strings.Repeat(" ", indent), format)
	return fmt.Printf(str, a...)
}

type KVTable struct {
	keys      []string
	values    []string
	maxKeyLen int
}

func NewKeyValueTable() (table *KVTable) {
	return &KVTable{}
}

func (t *KVTable) Add(key string, value string) {
	if len(key) > t.maxKeyLen {
		t.maxKeyLen = len(key)
	}
	t.keys = append(t.keys, key)
	t.values = append(t.values, value)
}

func (t *KVTable) Print(prefix string) {

	for i, key := range t.keys {
		pad := strings.Repeat(" ", t.maxKeyLen-len(key))
		fmt.Printf("%s%s:%s %s\n", prefix, key, pad, t.values[i])
	}

}
