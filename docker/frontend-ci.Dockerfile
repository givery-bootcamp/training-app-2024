FROM node:20-alpine
WORKDIR /usr/app
COPY . /usr/app/
RUN npm install && npm run build
EXPOSE 3000

ENTRYPOINT ["npm", "run", "start"]
