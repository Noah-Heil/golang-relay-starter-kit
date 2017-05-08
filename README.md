#Playing with Golang-React-Relay-Graphql

Dear Future Me (b/c God help your soul if you are not me and you are currently reading this/are going to try to use this!):

The entire application runs in docker containers because it is the best way to isolate env currently.

To run those containers I use docker-compose.

So to spin up "the stack" you need to `sudo su` and then `docker-compose up` using the -f command to specify the docker-compose files for
the env you plan on running it in.... Please note that the docker-compose files are currently specified with hard coded paths for my personal dev env.

Additionally the config/conf.toml file and anywhere else that I have specified configs for the postgres db are also hard coded dummy values....make sure to change them before going to prod.

Once you have spun up the stack you will need two more terminal windows one to run the cobra commands to spin up the graphql and graphiql servers and the other to run `npm start` from the root of the application to run the webpack build.

Currently everything runs but you need to work on getting the data flow from the db to the views fixed b/c it is not working b/c when you query in the graphiql ide you dont get the correct data...

so thats next step and once you are done with that you will have a working template for a golang application running all the tech you want!!!

Huge thanks to Hafiz Ismail who's graphql-go packages are awesome and I built off of and basically followed a ton of tutorials for...
here are some links:

[github profile](https://github.com/sogko)
[github graphql-go](https://github.com/graphql-go)
[golang-react-relay-graphql tut](https://wehavefaces.net/learn-golang-graphql-relay-2-a56cbcc3e341)
[golang-graphql first tut](https://wehavefaces.net/learn-golang-graphql-relay-1-e59ea174a902)

Soooo Hafiz if you ever see this...Your Awesome and thanks!