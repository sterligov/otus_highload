FROM node:14-alpine as builder

WORKDIR /dating

COPY *.json ./
RUN npm install
COPY ./ ${WORKDIR}
RUN npm run build

FROM nginx:1.17

COPY ./.docker/nginx.conf /etc/nginx/nginx.conf

RUN rm -rf /usr/share/nginx/html/*

COPY --from=builder /dating/build /usr/share/nginx/html

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]