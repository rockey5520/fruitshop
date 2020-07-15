FROM node:12-alpine


WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /app/package.json
COPY proxy.conf.json /app/proxy.conf.json

RUN npm install -g @angular/cli @angular-devkit/build-angular && npm install

EXPOSE 4200

# add app
COPY . /app

# start app
CMD npm start