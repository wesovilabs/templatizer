FROM scratch
COPY templatizer /
ENTRYPOINT ["/templatizer"]
