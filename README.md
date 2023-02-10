# caddy-imagefilter

根据 [caddy-imagefilter](https://github.com/ueffel/caddy-imagefilter) 魔改，增加了水印功能

### Examples

```caddy-d
:80 {
    @thumbnail {
        path_regexp thumb /w(400|800)(/.+\.(jpg|jpeg|png|gif|bmp|tif|tiff|webp))$
    }
    handle @thumbnail {
        rewrite {re.thumb.2}
        image_filter {
            watermark /path/to/image
        }
    }
}
```

