# Project Structure

```text
.
├── cmd
│   ├── cli
│   ├── cronjob         # for cronjob
│   └── server          # entry point, initialize the whole project
├── config              # contains config files, storing configurations
├── docs                # contains document files
├── global              # contains the global configurations for the whole project
├── internal            # all the logic, controllers, handlers, etc.
│   ├── controller      # contains controllers
│   ├── entity          # entities
│   ├── initialize      # initialize the application
│   │   ├── load_config # load the configuration into global setting in global
│   │   ├── logger      # start the global logger
│   │   ├── postgres    # connect to postgres
│   │   ├── redis       # connect to redis
│   │   ├── router      # initialize all routers, group routers
│   │   └── run         # The application actually starts running here
│   ├── middlewares     # middlewares
│   ├── model           # database model, similar to entities
│   ├── repo            # repositories
│   ├── routers         # routers for the app
│   ├── service         # services
│   └── wire            # set up wire dependency injection
├── log                 # can be used for storing log
├── migrations          # storing database migrations
├── pkg                 # code that can be shared among services
│   ├── logger          # constructor for global logger
│   ├── response        # common responses for API calls
│   │   ├── statusCode  # status code for internal use
│   │   ├── response    # common responses for API calls
│   │   └── settings    # config struct, used to read configuration into global
│   └── utils           # utilities
├── scripts             # scripts for setup, running tasks
├── storages            # storage
├── tests               # test files
└── third_party         # third party libraries
```
