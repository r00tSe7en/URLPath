# URLPath
批量处理url链接，获取多级路径并打印

链接文本
```
https://google.com/
https://example.com/api
http://example.com/books/all
https://example.com/books/all/1.pdf
http://noredirect.com/nor
https://redirect-no.fr/blocked
https://redirect-no.fr/Blocked
https://redirect-no.fr/Blocked/1/2/3/4/5
```
处理链接
```
cat url.txt|./URLPath
```
输出链接
```
https://google.com/
https://example.com/api
http://example.com/books
https://example.com/books
https://example.com/books/all
http://noredirect.com/nor
https://redirect-no.fr/blocked
https://redirect-no.fr/Blocked
https://redirect-no.fr/Blocked/1
https://redirect-no.fr/Blocked/1/2
https://redirect-no.fr/Blocked/1/2/3
https://redirect-no.fr/Blocked/1/2/3/4
```
# 参考
https://github.com/edoardottt/lit-bb-hack-tools/tree/main/cleanpath
