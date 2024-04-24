FROM node:21-alpine

WORKDIR /app

RUN npm install mathjs

COPY . .

CMD node --max-old-space-size=10000 /app/one.js && node --max-old-space-size=10000 /app/two.js && node --max-old-space-size=10000 /app/three.js