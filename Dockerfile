FROM alpine
ADD podApi /podApi

ENTRYPOINT ["/podApi"]
