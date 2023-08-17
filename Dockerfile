FROM scratch
ADD ./bin/main /bin/main

CMD ["/bin/main"]
