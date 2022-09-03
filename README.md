<h1 align="center">
  <img src="./docs/81841388.png" width="400" height='' alt="Noa Himesaka">
  <br>Noa Himesaka<br>  
</h1>
<h3 align="center">
    Download books from the <a href="https://book.sfacg.com/">sfacg</a> and 
    <a href="https://app.hbooker.com/">hbooker</a> to read them. 

</h3> 

[简中](./docs/README_zh-CN.md) | [繁中](./docs/README_zh-TW.md)

## **Functions**

- Download function is implemented for sfacg WeChat Api and hbooker Android API
- Login your account and save cookies to a ```config.json```
- Input the book id or url and download the book to the local directory
- Input url and download book text from the url
- Support download epub from sfacg and hbooker
- Search books by keyword,and download the search result
- [ **warning** ] New version book cache is incompatible with older book cache.

## **Example**

- --app=```<type[sfacg / cat]>```
- --account=```<account>```
- --password=```<password>```
- --download=```<type[bid / url]>```
- --search=```<keyword>```
- --show  < show the config.json file >

## **Disclaimers**

- This tool is for learning only. Please delete it from your computer within 24 hours after downloading.
- Please respect the copyright and do not spread the crawled books by yourself.
- In no event shall the authors or copyright holders be liable for any claim damages or other liability, whether in an
  action of contract tort or otherwise, arising from, out of or in connection with the software or the use or other
  dealings in the software , including but not limited to the use of the software for illegal purposes,author is not
  responsible for any legal consequences.


## **File tree**
``` 
C:.
│  .gitignore
│  config.json
│  go.mod
│  go.sum
│  LICENSE
│  main.go
│  README.md
│  
├─.idea
│      workspace.xml
│
├─cache
├─config
│      command.go
│      config.go
│      file.go
│      msg.go
│      thread.go
│      tool.go
│ 
├─docs
│      81841388.png
│      84782349.png
│      README_zh-CN.md
│      README_zh-TW.md
│
├─epub
│  │  dirinfo.go
│  │  epub.go
│  │  fetchmedia.go
│  │  fs.go
│  │  pkg.go
│  │  toc.go
│  │  write.go
│  │  xhtml.go
│  │
│  └─internal
│      └─storage
│          │  storage.go
│          │
│          ├─memory
│          │      file.go
│          │      fs.go
│          │
│          └─osfs
│                  fs.go
│
├─save
├─src
│  │  book.go
│  │  bookshelf.go
│  │  catalogue.go
│  │  login.go
│  │  progressbar.go
│  │  search.go
│  │
│  ├─boluobao
│  │      api.go
│  │
│  ├─hbooker
│  │  │  api.go
│  │  │  Geetest.go
│  │  │  UrlConstants.go
│  │  │
│  │  └─Encrypt
│  │          decode.go
│  │          Encrypt.go
│  │
│  └─https
│          Header.go
│          param.go
│          request.go
│          urlconstant.go
│
└─struct
    │  command.go
    │  config.go
    │
    ├─book_info
    │      book_info.go
    │
    ├─hbooker_structs
    │  │  chapter.go
    │  │  config.go
    │  │  content.go
    │  │  detail.go
    │  │  geetest.go
    │  │  key.go
    │  │  login.go
    │  │  recommend.go
    │  │  search.go
    │  │
    │  ├─bookshelf
    │  │      bookshelf.go
    │  │
    │  └─division
    │          division.go
    │
    └─sfacg_structs
        │  account.go
        │  book.go
        │  catalogue.go
        │  content.go
        │  login.go
        │  search.go
        │
        └─bookshelf
                bookshelf.go

```