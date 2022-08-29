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

Now, in order for _your project_ (example in the `project` folder) to talk to
that external project, you'll want to make some changes to your networking
section.

Reference the [project/docker-compose.yml](./project/docker-compose.yml) for
all changes, but you'll mostly notice that we're explicit by defining the
`default` network for things that your project defines, and talk to eachother,
and we explicitly define the `external` network for the things that need to
talk outward.
