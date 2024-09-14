FROM debian:stable-slim
COPY ./bin/rate-my-media-linux /bin/rate-my-media
CMD [ "/bin/rate-my-media" ]