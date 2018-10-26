---
title: Caching each instruction separately with asLayers
sidebar: reference
permalink: reference/build/as_layers.html
author: Alexey Igrychev <alexey.igrychev@flant.com>
summary: |
  <div class="language-yaml highlighter-rouge"><pre class="highlight"><code><span class="s">asLayers</span><span class="pi">:</span> <span class="s">true</span>
  </code></pre>
  </div>
---

User stages, `before_install`, `install`, `before_setup`, `setup` and `build_artifact`, depend upon the appropriate instructions in the configuration. Any modification in _stage_ instructions results in re-assembling the appropriate _stage_ with all instructions. Therefore, if the instructions are heavy and time-consuming, development of the configuration may take much time.

Let us also consider the situation where one of the last stage instruction fails. A user cannot retrieve the environment state preceding the failure of the instruction or check those previous instructions were correctly executed.

For development and debugging, we introduce _asLayers_ directive. When assembling, the instructions are cached separately, and re-assembly is only performed when their sequence changes. The directive can be specified for a specific _dimg_ or _artifact_ in the `dappfile.yaml` configuration.

If `asLayers: true`, then the new caching mode is enabled — one docker image for one shell command, or one task for ansible. Otherwise, if _asLayers_ directive is not specified (or if `asLayers: false`) default caching is applied, one docker image is used for all _stage_ instructions. 

Switching between assembly mode is only governed by the _asLayers directive_. Other instructions for the configuration remain unchanged. After having debugged the assembly instructions, _asLayers_ must be disabled.

_asLayers_ directive allows caching of individual instructions. If `--introspect-error` and `--introspect-before-error` introspection options are used, users may retrieve the environment before or after execution of a problem instruction.

> It is important not to use this instruction in the course of standard assembly of images: this mode generates an excessive number of docker images and is not intended for incremental assembly (due to a longer timeout and greater size of _stages cache_).