package password

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/howeyc/gopass"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Password() bool {
	passinit()
	verity := passverity()
	fmt.Println(verity)
	return verity
}

func passverity() bool {
	fmt.Print("请输入密码： ")
	inputpasswd, err := gopass.GetPasswd()
	if err != nil {
		panic(err)
	}
	sinputpasswd := sha256.New()
	sinputpasswd.Write([]byte(inputpasswd))
	sha256inputpasswd := hex.EncodeToString(sinputpasswd.Sum(nil))

	fmt.Println(sha256inputpasswd)

	filepass, err := os.Open("password.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer filepass.Close()
	savepasswd := make([]byte, 1024)
	upasswd := make([]byte, 1024)
	filepass.Read(savepasswd)
	for i, v := range savepasswd {
		if v == '$' {
			filepass.Seek(int64(i)+1, os.SEEK_SET)
			filepass.Read(upasswd)
		}
	}
	fmt.Println(string(upasswd))

	fmt.Println(sha256inputpasswd)
	fmt.Println(sha256inputpasswd == "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3")
	return false
}

func passinit() {
	if fileexists("password.txt") == true {
		return
	} else {
		createpasswd()
	}
}

func fileexists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

func createpasswd() {
	fmt.Print("请输入初始化密码： ")
	inputpasswd, err := gopass.GetPasswd()
	wsha256 := sha256.New()
	wsha256.Write([]byte(inputpasswd))
	psha256 := hex.EncodeToString(wsha256.Sum(nil))
	ssha256 := salt(5)
	wssha256 := sha256.New()
	wssha256.Write([]byte(ssha256))
	saltsha256 := hex.EncodeToString(wssha256.Sum(nil))
	file, err := os.OpenFile("password.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()
	file.Write([]byte(saltsha256))
	file.Write([]byte{'$'})
	file.Write([]byte(psha256))
	file.Sync()
}

func salt(num int) string {
	ss := make([]byte, 10)
	saltinit := []byte{'1', '2', '3', '4', '5', '6', 'a', 'b', 'c', 'd', 'e', 'f'}
	for i := 0; i < num; i++ {
		ss = append(ss, saltinit[rand.Intn(len(saltinit))])
	}
	return string(ss)
}
