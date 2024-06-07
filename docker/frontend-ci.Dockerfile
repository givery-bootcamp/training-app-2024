FROM node:20-alpine
WORKDIR /usr/app
COPY frontend /usr/app/
RUN ls -l
RUN npm install && npm run build
EXPOSE 3000

ENTRYPOINT ["npm", "run", "start"]
