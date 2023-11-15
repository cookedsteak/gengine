## gengine
本fork仓库，修复了原gengine出现的问题，以及添加了一些使用过程中需要的特性。
请随意取用，有问题可联系vx:cookedsteak

## 原项目文档
https://github.com/bilibili/gengine/wiki/

## 修改与添加的特性

- [x] 添加逻辑运算符兼容：'&','|','and','or'
- [x] 支持更高的精度类型(decimal)，并支持各种类型返回

#### more...

- [ ] 支持多元连续比较

## 一些学习资料
可参考：
https://tonybai.com/2022/05/25/an-example-of-implement-dsl-using-antlr-and-go-part2/

## antlr 的安装

下载jar包
```bash
$ export CLASSPATH=".:/usr/local/lib/antlr-4.13.1-complete.jar:$CLASSPATH"
$ alias antlr4='java -jar /usr/local/lib/antlr-4.13.1-complete.jar'
$ alias grun='java org.antlr.v4.gui.TestRig'
```

## 生产代码命令
cd 到g4文件所在目录
./antlr4.sh -Dlanguage=Go -encoding UTF-8 -visitor -o alr2 gengine.g4

## 调试
./antlr4.sh gengine.g4 生成java代码
./compile.sh ./gengine*.java 进行编译
./grun.sh gengine