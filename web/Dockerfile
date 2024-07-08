FROM nginx:alpine

WORKDIR /app

# Copy the built files from the previous stage
COPY .next /usr/share/nginx/html

# Copy custom NGINX configuration file
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80
EXPOSE 13081

# Start NGINX
CMD ["nginx", "-g", "daemon off;"]

# docker build -t cesslab/watchdog-ui:devnet .