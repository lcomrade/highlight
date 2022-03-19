[![Go report](https://goreportcard.com/badge/github.com/lcomrade/highlight)](https://goreportcard.com/report/github.com/lcomrade/highlight)
[![Go Reference](https://pkg.go.dev/badge/github.com/lcomrade/highlight.svg)](https://pkg.go.dev/github.com/lcomrade/highlight#section-documentation)
[![Release](https://img.shields.io/github/v/release/lcomrade/highlight)](https://github.com/lcomrade/highlight/releases/latest)
[![License](https://img.shields.io/github/license/lcomrade/highlight)](LICENSE)

**highlight** is a golang library for highlighting code syntax of different
programming and markup languages inside HTML documents.

## Install
Download from GitHub:
```
go get github.com/lcomrade/highlight
```

Or you can use a mirror:
```
go get git.lcomrade.su/root/highlight
```


## How it works?
The `<span>` tag is used to highlight syntax when converting to HTML.
This tag does not do anything on its own,
but you can assign different properties to it in CSS.
Including color and font.

Before highlighting:
```robots.txt
User-agent: * # comment
Disallow: /faq

# comment 1
# comment 2

Allow: /
Crawl-delay: 10
Sitemap: https://example.org/sitemap.xml
Host: https://mirror.example.org
```

After highlighting:
```html
<span class='code-keyword'>User-agent:</span> * <span class='code-c'># comment</span>
<span class='code-keyword'>Disallow:</span> /faq

<span class='code-comment'># comment 1</span>
<span class='code-comment'># comment 2</span>

<span class='code-keyword'>Allow:</span> /
<span class='code-keyword'>Crawl-delay:</span> 10
<span class='code-keyword'>Sitemap:</span> https://example.org/sitemap.xml
<span class='code-keyword'>Host:</span> https://mirror.example.org
```


## Documentation
- Offline documentation: `go doc -all github.com/lcomrade/highlight`
- [Online documentation](https://pkg.go.dev/github.com/lcomrade/highlight#section-documentation)
- [Changelog](CHANGELOG.md)
