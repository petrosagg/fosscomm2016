FROM resin/armhf-alpine:3.3

COPY qemu-arm-static /usr/bin/qemu-arm

COPY sh-shim /bin/sh.shim

RUN [ "qemu-arm", "/bin/cp", "-f", "/bin/sh", "/bin/sh.real" ]

RUN [ "qemu-arm", "/bin/ln", "-f", "/bin/sh.shim", "/bin/sh" ]

RUN echo foobar
