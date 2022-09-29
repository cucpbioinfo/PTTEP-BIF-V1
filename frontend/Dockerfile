# base image
FROM node:15.2.1-alpine

# working directory
WORKDIR /app

# install and cache app dependencies
COPY package.json .
COPY yarn.lock .
RUN yarn install

# copy app files and build
COPY . .
RUN yarn build

EXPOSE 3000

# start app
CMD [ "yarn", "start" ]