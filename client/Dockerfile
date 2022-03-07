FROM node:14.9

WORKDIR /017
COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3000

VOLUME [ "/017" ]

CMD ["npm", "start"]