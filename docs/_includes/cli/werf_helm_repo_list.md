{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}
List chart repositories

{{ header }} Syntax

```bash
werf helm repo list [flags] [options]
```

{{ header }} Options

```bash
      --debug=false:
            enable verbose output. Defaults to $HELM_DEBUG
      --helm-home='/home/aigrychev/.helm':
            location of your Helm config. Defaults to $HELM_HOME
  -h, --help=false:
            help for list
```

