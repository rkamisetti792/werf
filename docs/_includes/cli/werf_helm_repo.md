{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}

This command consists of multiple subcommands to interact with chart repositories.
It can be used to init, add, remove, and list chart repositories


{{ header }} Options

```bash
  -h, --help=false:
            help for repo
```

