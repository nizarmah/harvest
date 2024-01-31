# use official nginx image
FROM nginx:alpine

# set working dir
WORKDIR /usr/share/nginx

# remove default nginx html files
RUN rm -rf ./html/*

# copy pear code
COPY pear/ ./html/

# expose port 80
EXPOSE 80

# start nginx
CMD ["nginx", "-g", "daemon off;"]
