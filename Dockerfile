FROM node:15.5.1-alpine3.10

# install simple http server for serving static content
RUN npm install -g http-server

# make the 'app' folder the current working directory
WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY frontend/package*.json ./

# install project dependencies
RUN npm install

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY frontend .

# build app for production with minification
RUN npm run build

EXPOSE 8080
CMD [ "npm", "run server" ]