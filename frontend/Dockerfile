
FROM node:latest AS builder

WORKDIR /app

COPY . .

# Set BASE_URL during build
ARG BASE_URL
ENV BASE_URL=${BASE_URL}

RUN npm install && npm run build


FROM nginx:alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY --from=builder /app/dist/spa .

ENTRYPOINT ["nginx", "-g", "daemon off;"]
