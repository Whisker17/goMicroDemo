FROM alpine
ADD test2-srv /test2-srv
ENTRYPOINT [ "/test2-srv" ]
