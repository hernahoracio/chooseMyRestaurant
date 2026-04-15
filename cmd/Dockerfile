ARG BUILD_IMAGE=chainguard/go:latest
ARG RUNTIME_IMAGE=chainguard/static:latest

# ============================
# Installer stage (Downloads the Go modules)
# ============================
FROM ${BUILD_IMAGE} AS installer

WORKDIR /src

COPY go.mod go.sum pkg/model/go.mod pkg/model/go.sum ./

RUN  go mod download

# ============================
# Builder stage (compiles the Go binary)
# ============================
FROM installer AS builder

WORKDIR /src

COPY . .

RUN CGO_ENABLED=1 go build -o /out/cmr ./cmd/cmr

# ============================
# RunTime stage (Executes the Go binary as a non-root user)
# ============================
FROM ${RUNTIME_IMAGE} AS runtime

COPY --from=builder /out/cmr /cmr

EXPOSE 8085

CMD ["/cmr"]
