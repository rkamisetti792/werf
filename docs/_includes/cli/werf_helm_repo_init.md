{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}
Init chart repositories

{{ header }} Syntax

```bash
werf helm repo init [flags] [options]
```

{{ header }} Options

```bash
      --debug=false:
            enable verbose output. Defaults to $HELM_DEBUG
      --helm-home='/home/aigrychev/.helm':
            location of your Helm config. Defaults to $HELM_HOME
  -h, --help=false:
            help for init
      --local-repo-url='http://127.0.0.1:8879/charts':
            URL for local repository
      --skip-refresh=false:
            do not refresh (download) the local repository cache
      --stable-repo-url='https://kubernetes-charts.storage.googleapis.com':
            URL for stable repository
```

