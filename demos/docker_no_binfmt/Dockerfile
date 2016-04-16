FROM resin/armhf-alpine:3.3

COPY qemu-arm-static /usr/bin/qemu-arm

RUN [ "qemu-arm", "/bin/echo", "Hello from ARM Docker container :)" ]
