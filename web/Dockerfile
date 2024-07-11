FROM node:20-alpine
WORKDIR /app
ENV NODE_ENV production

COPY package.json package-lock.json ./
COPY public ./public
COPY .next/standalone ./
COPY .next/static ./.next/static
COPY .next/server ./.next/server

EXPOSE 13080
ENV PORT 13080
ENV HOSTNAME "0.0.0.0"
CMD node server.js