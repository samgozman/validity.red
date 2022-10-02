FROM gcr.io/distroless/static-debian11
COPY userApp .
CMD ["/userApp"]
