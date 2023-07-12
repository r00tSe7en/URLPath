# URLPath

![URLPath](https://socialify.git.ci/r00tSe7en/URLPath/image?description=1&forks=1&issues=1&language=1&name=1&owner=1&pattern=Floating%20Cogs&pulls=1&stargazers=1&theme=Dark)

# 1.配合爬虫使用

这里联动projectdiscovery的katana爬虫，拿特斯拉官网测试，最终结果是1441行并无重复

https://github.com/projectdiscovery/katana
```
katana -u https://tesla.com | ./URLPath | tee result.txt
...
   ...
   
wc -l result.txt
1441 //统计行数

cat result.txt | uniq | wc -l
1441 //uniq去重后统计行数
```

# 2.处理链接文本

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
https://example.com/?id=1
https://example.com/test/?id=1
```
处理链接
```
Take as input on stdin a list of urls and print on stdout all the unique paths (at any level).
        $> cat input | URLPath -l 2 
        -l x is specify number of path levels to print. -1 or no parameters specified to print all path levels.
         

1. cat url.txt|./URLPath  //打印全部级别路径

2. cat url.txt|./URLPath -l 2  //打印2级及2级以下级别路径
```
输出链接
```
(提示：无path的url不会打印)
http://example.com/api
http://example.com/Api
https://example.com/books
https://example.com/books/all
https://example.com/1
https://example.com/1/2
https://example.com/1/2/3
https://example.com/1/2/3/4
https://example.com/1/2/3/4/5
https://example.com/test
```

# 参考
https://github.com/edoardottt/lit-bb-hack-tools/tree/main/cleanpath
