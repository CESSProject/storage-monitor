FROM node:20-alpine
WORKDIR /app
ENV NODE_ENV production

COPY package.json package-lock.json ./
COPY public ./public
COPY .next/standalone ./
COPY .next/static ./.next/static
COPY .next/standalone/node_modules ./node_modules

EXPOSE 13080
ENV PORT 13080
CMD HOSTNAME="0.0.0.0" node server.js