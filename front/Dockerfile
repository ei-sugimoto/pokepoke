FROM node:bullseye as base

# development stage
FROM base as development
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
CMD ["npm", "run", "dev"]