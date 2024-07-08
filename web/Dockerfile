FROM node:20.11.0 AS base

WORKDIR /app

ARG NEXT_PUBLIC_API_SERVER_URL
ENV NEXT_PUBLIC_API_SERVER_URL=${NEXT_PUBLIC_API_SERVER_URL}

COPY package.json package-lock.json* ./

RUN npm ci

COPY . .

RUN npm run build

FROM base AS runner
WORKDIR /app

ENV NODE_ENV production

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./

EXPOSE 13081

CMD ["npm", "run", "start",  "--",  "--port", "13081"]

# docker build --build-arg NEXT_PUBLIC_API_SERVER_URL=http://watchdog:13080 -t watchdog-web .