# go run fileName
除了使用go build命令外，Go语言还为我们提供了go run命令，go run命令将编译和执行指令合二为一，会在编译之后立即执行Go语言程序，但是不会生成可执行文件。

# go build fileName 
1. 当参数不为空时
如果 fileName 为同一 main 包下的所有源文件名（可能有一个或者多个），编译器将生成一个与第一个 fileName 同名的可执行文件（如执行go build abc.go def.go ...会生成一个 abc.exe 文件）；如果 fileName 为非 main 包下的源文件名，编译器将只对该包进行语法检查，不生成可执行文件。

2. 当参数为空时
如果当前目录下存在 main 包，则会生成一个与当前目录名同名的“目录名.exe”可执行文件（如在 hello 目录中执行go build命令时，会生成 hello.exe 文件）；如果不存在 main 包，则只对当前目录下的程序源码进行语法检查，不会生成可执行文件

