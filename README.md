# okproxy
使用goproxy实现go module的简单包管理 方便搭建私服和离线使用

# Quick start


1. go env -w GO111MODULE=on
2. go env -w GOPROXY=https://goproxy.cn,direct
3. go get github.com/dalixu/okproxy
4. run ./okproxy
5. go env -w GOPROXY=http://localhost:8899,direct
6. go get ... 
