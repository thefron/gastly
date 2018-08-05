FROM node:10.8

RUN yarn global add @vue/cli

WORKDIR /app

COPY package.json yarn.lock /app/

RUN yarn install

COPY . /app/

EXPOSE 8080
CMD ["yarn", "serve"]
