## gengine
本fork仓库用于修复线上gengine出现的问题

## 原项目文档
https://github.com/bilibili/gengine/wiki/

## 通过ANTLR修改
可参考：
https://tonybai.com/2022/05/25/an-example-of-implement-dsl-using-antlr-and-go-part2/

## antlr installation
export CLASSPATH=".:/usr/local/lib/antlr-4.13.1-complete.jar:$CLASSPATH"
alias antlr4='java -jar /usr/local/lib/antlr-4.13.1-complete.jar'
alias grun='java org.antlr.v4.gui.TestRig'

## generator

antlr4 -Dlanguage=Go -visitor -o alr2 gengine.g4 