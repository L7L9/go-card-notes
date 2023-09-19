package utils

import (
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"io"
	"lqlzzz/go-card-notes/global"
	"mime/multipart"
)

var ipfsClient *shell.Shell

// InitIpfsClient //
// 初始化ipfs
func InitIpfsClient() {
	addr := fmt.Sprintf("%s:%d", global.GCN_CONFIG.Ipfs.Host, global.GCN_CONFIG.Ipfs.Port)
	ipfsClient = shell.NewShell(addr)
}

// IpfsAdd //
// 向ipfs中添加文件
func IpfsAdd(file multipart.File) (string, error) {
	cid, err := ipfsClient.Add(file)
	if err != nil {
		return "", err
	}
	return cid, nil
}

// IpfsGet //
// 向ipfs中获取
func IpfsGet(cid string) (io.ReadCloser, error) {
	readCloser, err := ipfsClient.Cat(cid)
	if err != nil {
		return nil, err
	}
	return readCloser, nil
}
