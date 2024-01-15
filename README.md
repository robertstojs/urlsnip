# urlsnip

<img src="./img/logo.svg" align="right"
     alt="urlsnip logo by Robert Štojs" width="192" height="192">

urlsnip is a tiny URL shortening and redirection program written in Go. It intercepts web requests and instantly redirects them according to a list of regular expression rules.

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
2024/01/15 14:07:24 Redirected via regex 'blog1' to 'https://example.com/blog-page' using pattern '^blog[0-9]+$'
2024/01/15 14:07:30 Redirected via regex 'blog2' to 'https://example.com/blog-page' using pattern '^blog[0-9]+$'
```

```shell
git clone https://github.com/robertstojs/urlsnip.git
cd urlsnip
go build
```

















