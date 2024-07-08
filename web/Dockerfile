FROM node:20.11.0

WORKDIR /app

ENV NODE_ENV production
ENV NEXT_PUBLIC_API_SERVER_URL=http://cess-monitor:13080

COPY package.json package-lock.json* ./
COPY public ./public
COPY .next ./.next

EXPOSE 13081

CMD ["npm", "run", "start",  "--",  "--port", "13081"]
# docker build -t cesslab/watchdog-web .