# App Management for K8s

I've played with things like Helm, and I'm keen to try ksonnet and
draft. My attempts will end up in this repo.

## Old School

After a few evolutions, the solution I've been most fond of, so far, is
Jinja2 templating for managing resources that are deployed to all
environments, with a `app.yaml` file that stores the defaults and
differences in each environment, when it comes to replicas, and
environment variables.

A Jinja2 CLI that supports YAML "context" for the templates was a great
solution for this.
