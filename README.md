# caddy-imagefilter

根据 [caddy-imagefilter](https://github.com/ueffel/caddy-imagefilter) 魔改，增加了水印功能

## 安装
```bash
xcaddy build --with github.com/hakutyou/caddy-imagefilter/watermark
```

## 使用
```caddy-d
{
    order image_filter before file_server
}

your-web.site {
    root * /path/to/website

    file_server

    # 所有的图片
    @thumbnail {
        path_regexp thumb /.+\.(jpg|jpeg|png|gif|bmp|tif|tiff|webp)$
    }

    image_filter @thumbnail {
        # 只支持 png 图片
        watermark /path/to/image.png
    }
}
```

