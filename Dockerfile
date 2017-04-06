FROM floreks/kubepi-base

COPY build/sensor-mqtt-client-arm /usr/bin/sensor-mqtt-client-arm

ENTRYPOINT ["/usr/bin/sensor-mqtt-client-arm"]