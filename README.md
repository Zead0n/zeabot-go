# NOTICE
This project is going through a major rework.

# zeabot-go
A recreation of [zeabot-rs](https://github.com/Zead0n/zeabot-rs) written in go.

A Discord bot that uses [Disgo](https://github.com/DisgoOrg/disgo), [Disgolink](https://github.com/disgoorg/disgolink), and [Lavalink](https://github.com/lavalink-devs/Lavalink) to play audio from a youtube source (possibly others) in a voice chat.

# How to use
It is hightly recommended to run this with [docker compose](https://docs.docker.com/compose/) as it will run the bot and the Lavalink container together.

Here is the basic `docker-compose.yml` contents:
```yaml
services:
  bot:
    image: zead0n/zeabot-go:latest
    container_name: zeabot-go
    restart: unless-stopped
    depends_on:
      - lavalink
    environment:
      DISCORD_TOKEN: "TOKEN"
      LAVALINK_HOSTNAME: "lavalink"
      LAVALINK_PORT: "2333"
      LAVALINK_PASSWORD: "CHANGE_ME"
    networks:
      - zeabot

  lavalink:
    image: ghcr.io/lavalink-devs/lavalink:latest-alpine 
    container_name: lavalink
    restart: unless-stopped
    environment:
      _JAVA_OPTIONS: "-Xmx1G"

      # Server config
      SERVER_PORT: "2333"
      SERVER_ADDRESS: "0.0.0.0"
      LAVALINK_SERVER_PASSWORD: "CHANGE_ME"
      SERVER_HTTP2_ENABLED: "false"

      # youtube-source plugin
      # Replace 'VERSION' with the actual version you wish to use. refer to https://github.com/lavalink-devs/youtube-source/releases
      LAVALINK_PLUGINS_0_DEPENDENCY: "dev.lavalink.youtube:youtube-plugin:VERSION"
      LAVALINK_PLUGINS_0_REPOSITORY: "https://maven.lavalink.dev/releases"
      LAVALINK_PLUGINS_0_SNAPSHOT: "false"
      PLUGINS_YOUTUBE_ENABLED: "true"
      PLUGINS_YOUTUBE_ALLOWSEARCH: "true"

      # disable default youtube source
      LAVALINK_SERVER_SOURCES_YOUTUBE: "false"
    networks:
      - zeabot

networks:
  zeabot:
    name: zeabot
```
Then run `docker compose up -d`.
