FROM gcr.io/distroless/static-debian11
COPY brokerApp .
CMD ["/brokerApp"]
