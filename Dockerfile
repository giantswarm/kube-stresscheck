FROM alpine:3.7

ADD kube-stresscheck /usr/bin/kube-stresscheck
ENTRYPOINT ["/usr/bin/kube-stresscheck"]
