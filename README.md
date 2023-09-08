## gengine
本fork仓库，修复了原gengine出现的问题，以及添加了一些使用过程中需要的特性。
请随意取用，有问题可联系vx:cookedsteak

## 原项目文档
https://github.com/bilibili/gengine/wiki/

## 修改与添加的特性

- [x] 添加逻辑运算符兼容：'&','|','and','or'
- [x] 添加比较运算符兼容：'lt','gt','ge','le','eq','ne'
- [x] 支持更高的精度类型(decimal)，并返回string类型

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
antlr4 -Dlanguage=Go -visitor -o alr2 gengine.g4 