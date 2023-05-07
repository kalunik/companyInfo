FROM alpinelinux/golang
LABEL authors="kalunik"

ENTRYPOINT ["top", "-b"]