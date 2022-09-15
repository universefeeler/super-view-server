package example

import (
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"io/fs"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestParseForCidFromUrl(t *testing.T) {
	url1 := "https://bafkreih6mvlagcf5u4elijd72v7xn6oi63sdnodo3swjledwzh4d3v4efe.ipfs.dweb.link"
	hashArray := strings.Split(url1, "/")
	cid := hashArray[len(hashArray)-1]
	if len(cid) < 45 {
		fmt.Println()
	}
	cid = strings.Split(cid, ".")[0]
	fmt.Println(cid)

}

func TestIpfsWithNode(t *testing.T) {
	ipfs1 := "QmbC6ir2SohLkvZ2ycvYZp2kg6D9A5gR56JTXzyzrEmMuo"
	fmt.Println(len(ipfs1))
	//CatIPFS(ipfs1)
	ipfs2 := "bafybeiehvq5gbntfb7snqlwtdbxtihc4qa3eozzgsduazv66aoioohpeqq"
	fmt.Println(len(ipfs2))
	//CatIPFS(ipfs2)
	ipfs3 := "QmaC4dhAKB4sWR319nzStRHzzr6JAvWrrRTCPdZxTyFbas"
	GetIPFS(ipfs3)
}

func TestIpfsWithHttp(t *testing.T) {
	ipfs1 := "https://ipfs.io/ipfs/bafybeiehvq5gbntfb7snqlwtdbxtihc4qa3eozzgsduazv66aoioohpeqq"
	req, _ := http.NewRequest("GET", ipfs1, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	//ipfs3 := "QmeZjspAC73E9nsfmGXP7DaCemkVfHsKBosE44tkcKp4ma"
}

func GetIPFS(hash string) {
	//sh := shell.NewShell("localhost:5001")
	sh := shell.NewShell("https://ipfs.io")
	object, err := sh.ObjectGet(hash)
	fmt.Println(object.Links, err)
}

func CatIPFS(hash string) {
	//sh := shell.NewShell("localhost:5001")
	sh := shell.NewShell("https://ipfs.io")
	read, err := sh.Cat(hash)
	if err == nil {
		body, err2 := ioutil.ReadAll(read)
		fmt.Println(err2)
		ToFileWithBytes(body, hash+".png")
	}
}

func ToFileWithBytes(fileBytes []byte, suffix string) {
	err3 := ioutil.WriteFile(suffix, fileBytes, fs.ModePerm)
	fmt.Println(err3)
}
