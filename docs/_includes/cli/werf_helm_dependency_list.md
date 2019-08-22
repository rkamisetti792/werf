{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}

List all of the dependencies declared in a chart.

This can take chart archives and chart directories as input. It will not alter
the contents of a chart.

This will produce an error if the chart cannot be loaded. It will emit a warning
if it cannot find a requirements.yaml.


{{ header }} Syntax

```bash
werf helm dependency list [flags] [options]
```

{{ header }} Options

```bash
      --debug=false:
            enable verbose output. Defaults to $HELM_DEBUG
      --dir='':
            Change to the specified directory to find werf.yaml config
      --helm-home='/home/aigrychev/.helm':
            location of your Helm config. Defaults to $HELM_HOME
  -h, --help=false:
            help for list
```

