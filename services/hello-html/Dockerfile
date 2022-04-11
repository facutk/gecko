FROM node:lts-alpine as builder

WORKDIR /usr/src/app

COPY package*.json /usr/src/app

RUN npm install

COPY ./ /usr/src/app

RUN npm run build


FROM nginx:alpine

EXPOSE 80

# COPY ./index.html /usr/share/nginx/html/index.html

COPY --from=builder /usr/src/app/build/ /usr/share/nginx/html
COPY ./etc/nginx.conf /etc/nginx/nginx.conf
