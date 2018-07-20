FROM alpine:3.7

# Build stress from source.
RUN STRESS_VERSION=1.0.4; \
    apk add --no-cache g++ make curl && \
    curl https://people.seas.harvard.edu/~apw/stress/stress-${STRESS_VERSION}.tar.gz | tar xz && \
    cd stress-${STRESS_VERSION} && \
    ./configure && make && make install && \
    apk del --purge g++ make curl && rm -rf stress-*

ADD kube-stresscheck /usr/bin/kube-stresscheck
ENTRYPOINT ["/usr/bin/kube-stresscheck"]
