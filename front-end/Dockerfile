FROM scratch
MAINTAINER David Tennander <davidten@kth.se>
ADD bin/server /server
ADD login-page.tmpl /gateway/login-page.tmpl
ENTRYPOINT ["/server"]