# Example of Docker Compose (and working together)

Ever had a compose project and wanted to be able to talk to it locally using
your compose, yet still separately?

This shows you how to do that.

## Setup

First, get into the `external` project. We'll use this to showcase an example
other project that you don't want to change or affect.. just use it.

```
docker compose up -d
```

### Changes to your project

Now, in order to talk to that project, you're going to need to make some
network changes to yours.. You can see those changes in the `project` folder
under its `docker-compose.yml`
