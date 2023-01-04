# URLPath
批量处理url链接，获取多级路径并打印

链接文本
```
http://example.com
http://example.com/api
http://example.com/Api
https://example.com/books/all
https://example.com/1.pdf
https://example.com/books/all/1.pdf
https://example.com/1/2/3
https://example.com/1/2/3/4/5
```
处理链接
```
cat url.txt|./URLPath
```
输出链接
```
http://example.com
http://example.com/api
http://example.com/Api
https://example.com/books
https://example.com/books/all
https://example.com/1
https://example.com/1/2
https://example.com/1/2/3
https://example.com/1/2/3/4
https://example.com/1/2/3/4/5
```
# 参考
https://github.com/edoardottt/lit-bb-hack-tools/tree/main/cleanpath
