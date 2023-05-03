# ~35mb final image size

# Build backend binary file
FROM golang:1.16-alpine as backend_build

COPY backend /go/src/app/backend

WORKDIR /go/src/app/backend
RUN apk add build-base
RUN go build .

# Build static frontend files
# FROM node:18-alpine3.14 as frontend_build

# COPY frontend /app/frontend

# ENV PATH /app/node_modules/.bin:$PATH

# WORKDIR /app/frontend

# RUN npm ci
# RUN npm run build

# Slim production container with go binary and static files
# FROM alpine

ARG BUILD_ENV
ENV BUILD ${BUILD_ENV}
ENV CF_ACCOUNTS cf_accounts
ENV ADMIN_TOKEN admin_token
ENV DATABASE_URL localhost
ENV ADMIN_ACCOUNTS []
ENV PORT 8000

EXPOSE ${PORT}

# COPY --from=frontend_build /app/frontend/build /app/frontend/build
# COPY --from=backend_build /go/src/app/backend/servermodule /app/backend/servermodule
# COPY --from=backend_build /go/src/app/backend/configs/data /app/backend/configs/data

# WORKDIR /app/backend

CMD ["./servermodule"]
