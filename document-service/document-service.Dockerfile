FROM gcr.io/distroless/static-debian11
COPY documentApp .
CMD ["/documentApp"]
