FROM node:23.9-alpine AS build

WORKDIR /app
COPY package*.json ./
RUN "npm install"
COPY . .
RUN "npm run build"

FROM nginx:stable-alpine AS production
COPY --from=build /app/build /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]