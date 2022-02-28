FROM alpine:3.7
ADD bin/gh-templatizer.linux /usr/local/bin/gh-templatizer
ENTRYPOINT ["gh-templatizer"]