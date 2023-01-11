FROM node:19

WORKDIR /usr/src/app/client/Vue3

COPY ./Vue3/package*.json ./
RUN npm install
COPY ./Vue3 .
RUN npm run build

CMD  ["npm", "run", "serve"]