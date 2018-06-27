# vote-cli

提供数据库清空接口

## 如何跑起来

1. 抓取项目

  ````
  go get github.com/Actoors/vote-cli
  ````

2. 解决依赖

   ```
   cd $GOPATH/src/github.com/Actoors/vote-cli
   dep ensure
   ```

3. 将`conf_reference.json`复制到`conf.json`，修改`conf.json`中数据源配置

4. build并运行

   ```
   go build
   ./vote-cli
   ```

   或者使用docker

   ```
   docker build -t vote-cli .
   docker run vote-cli
   ```