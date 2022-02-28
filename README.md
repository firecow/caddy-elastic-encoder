# caddy-elastic-encoder
Caddy Logging Encoder that follows ECS guidelines

https://www.elastic.co/guide/en/ecs/current/index.html

[![Github Action](https://github.com/firecow/caddy-elastic-encoder/actions/workflows/go-qa.yml/badge.svg)](https://github.com/firecow/caddy-elastic-encoder/actions/workflows/go-qa.yml)
[![Release](https://img.shields.io/github/v/release/firecow/caddy-elastic-encoder?sort=semver)](https://github.com/firecow/caddy-elastic-encoder)
[![License](https://img.shields.io/github/license/firecow/gitlab-ci-local)](https://github.com/firecow/caddy-elastic-encoder)
[![Renovate](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://renovatebot.com)

```caddyfile
{
    log {
        level info
        format elastic
    }
}

:80
reverse_proxy http://webserver:3000
```