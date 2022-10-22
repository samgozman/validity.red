FROM gcr.io/distroless/static-debian11
COPY gatewayApp .
CMD ["/gatewayApp"]
