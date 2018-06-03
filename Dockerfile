FROM alpine:latest

COPY vim-tips-web /go/bin/vim-tips/
COPY templates /go/bin/vim-tips/templates
COPY screenshots /go/bin/vim-tips/screenshots
COPY public /go/bin/vim-tips/public

WORKDIR /go/bin/vim-tips

CMD ["/go/bin/vim-tips/vim-tips-web"]
EXPOSE 3000