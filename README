==========
 Buildenv
==========

Buildenv is a simple tool to perform a series of steps based on a YAML
file. The reason for buildenv is because we use Jenkins and the flow
groovy DSL to perform some steps. This works well, but it requires you
run the flow within Jenkins. Buildenv aims to provide similar
semantics to run the steps on your local machine.


Usage
=====

The essence of buildenv is the steps YAML file. Here is an example

```yaml

---
- Name: Do it
  Command: echo "Doing it!"

- Name: build backend
  Parallel: true
  Steps:

    - Name: build foo
      Command: echo "make foo"

    - Name: build bar
      Command: echo "make bar"

    - Name: build baz
      Command: echo "make baz"

- Name: Wrap up
  Command: echo "uploading artifacts"
```

You can then run the steps via `buildenv --steps steps.yml`. It will
run the parallel steps in parallel and wait before finishing up any
other steps.

Buildenv is focused on calling commands, therefore if you wanted to
distribute the tasks over a cluster of machine (like Jenkins flow),
you can build that into your commands. For example, if you configured
your environment to use a docker swarm, you could use `docker run`
commands on the swarm and utilize it as a cluster.

Status
======

Buildenv essentially works, but there are somethings I'd like to add:

 - retries
 - configure stop or continue on failure
 - colored output to differentiate between the different steps
 - better logging / reporting
