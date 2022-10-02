FROM gcr.io/distroless/static-debian11

COPY loggerApp .
CMD ["/loggerApp"]
