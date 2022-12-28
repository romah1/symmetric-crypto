# Simple Symmetric Crypto example

The program takes 3 arguments:
1. `--file` is file that contains a list of strings
2. `--numbilets` is maximum number of tickets
3. `--parameter` is used for changing the distribution (like seed)
## Example input
```
Hello Worldovich
World Hellovich
Evil Arthas
Dunder Mifflin This is Pam
```
## How to run
```
go build main.go &&  ./main --file input --numbilets 7 --parameter 3
```
## Example output
```
Hello Worldovich: 7
World Hellovich: 5
Evil Arthas: 2
Dunder Mifflin This is Pam: 1
```