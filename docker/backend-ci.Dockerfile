FROM ubuntu:22.04
WORKDIR /app
COPY backend/myapp /app
RUN chmod +x myapp

ENTRYPOINT ["./myapp"]
