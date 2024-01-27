# urlsnip
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![Build](https://github.com/robertstojs/urlsnip/actions/workflows/build.yml/badge.svg)

<img src="./img/logo.svg" align="right" width="156px">

urlsnip is a tiny URL shortener and routing program written in Go. It intercepts web requests and instantly redirects them according to a list of regular expression rules.

* Simple to use.
* **Regular expression** support.
* Uses **standard libraries**.
* Configurable with **json**.
* Logs natively to **syslog**.

```shell
# port defaults to 8080 if empty
./urlsnip --config=./config.json --port=8090
2024/01/15 14:07:11 Configuration file loaded successfully
2024/01/15 14:07:11 Server starting on :8090
2024/01/15 14:07:21 No redirect mapping found for url 'testing'
2024/01/15 14:07:24 Redirected via route 'blog1' to 'https://example.com/blog-page' using pattern '^blog[0-9]+$'
2024/01/15 14:07:30 Redirected via route 'blog2' to 'https://example.com/blog-page' using pattern '^blog[0-9]+$'
```

```shell
git clone https://github.com/robertstojs/urlsnip.git
cd urlsnip
go build
```
