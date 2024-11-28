# Ledger - A job task

## Description
This repository contains a simple leader transaction record system. 

## Installation
To run this project, you must have docker installed. Rest of the things are taken care of with `docker compose`. 

Once docker is setup, run:

```bash 
docker compose up -d
```

## Features
I built a custom opinionated and minimalistic framework / engine that is also used in this project. The internal working can be found under `Engine` package. 

It features glue components together for rest server, apply middleware globally or per route basis and use no external dependency. I plan to implement gRPC within the engine as well.

## Contact
Please feel free to reach out to me for any feedback or questions. Thank you.
