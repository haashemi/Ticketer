FROM node:18-alpine AS builder
WORKDIR /app

# Enable PNPM
RUN corepack enable

# Copy and install dependencies
COPY package.json .
COPY pnpm-lock.yaml .
RUN pnpm install

# Copy the whole source code and build the project
COPY . .

ARG API_HOST
ENV API_HOST=$API_HOST

RUN pnpm check
RUN pnpm run build
RUN pnpm prune --production

FROM node:18-alpine
WORKDIR /app

COPY --from=builder /app/build build/
COPY --from=builder /app/node_modules node_modules/
COPY package.json .

EXPOSE 3000
ENV NODE_ENV=production
CMD [ "node", "build" ]