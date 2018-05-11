# base image
FROM node:latest

# add `/usr/src/app/node_modules/.bin` to $PATH
ENV PATH /hatch/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /hatch/package.json
# set working directory
WORKDIR /hatch

RUN yarn

# start app
CMD ["yarn", "run", "watch"]