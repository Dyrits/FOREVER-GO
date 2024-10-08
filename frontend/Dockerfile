FROM node:lts-alpine

WORKDIR /app

COPY package.json ./

RUN npm install

COPY ./ ./

ENV NODE_OPTIONS=--openssl-legacy-provider

CMD ["npm", "start"]