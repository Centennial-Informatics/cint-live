FROM node:16-alpine

ENV PORT 3000
EXPOSE ${PORT}

WORKDIR /app

COPY . .

RUN npm ci

CMD ["npm", "run", "dev", "--host"]
