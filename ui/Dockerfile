FROM cypress/included:4.5.0
# replace this with your application's default port
WORKDIR /work
COPY yarn.lock .
COPY package.json .
RUN yarn

COPY . .
ENV CYPRESS_BASE_URL=http://localhost:3001
ENV REACT_APP_GQL_URL=http://food:4242/query
EXPOSE 3001
RUN pwd
RUN ls
ENTRYPOINT [ "./ci.sh" ]