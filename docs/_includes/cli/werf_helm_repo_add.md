{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}
Add a chart repository

{{ header }} Syntax

```bash
werf helm repo add [flags] [NAME] [URL] [options]
```

{{ header }} Options

```bash
      --ca-file='':
            verify certificates of HTTPS-enabled servers using this CA bundle
      --cert-file='':
            identify HTTPS client using this SSL certificate file
      --debug=false:
            enable verbose output. Defaults to $HELM_DEBUG
      --helm-home='/home/aigrychev/.helm':
            location of your Helm config. Defaults to $HELM_HOME
  -h, --help=false:
            help for add
      --key-file='':
            identify HTTPS client using this SSL key file
      --no-update=false:
            raise error if repo is already registered
      --password='':
            chart repository password
      --username='':
            chart repository username
```

