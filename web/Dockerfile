FROM nginx:alpine

COPY out /usr/share/nginx/html

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 13081

# 启动 Nginx
CMD ["nginx", "-g", "daemon off;"]
