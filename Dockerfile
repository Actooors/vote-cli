FROM ubuntu:14.04
MAINTAINER mzz m@mzz.pub
ADD vote-cli ./
ADD conf.json conf.json
EXPOSE 8128
ENTRYPOINT ["vote-cli"]