# ScourgeBuffBot
## Discord Bot

![alt text](https://i.imgur.com/ptoAn69.png)

This is a bot to copy carlbot embeds from a server and paste them on another.
Works with multiple dicord servers and includes some bonus features
Written in Go with discordgo (https://github.com/bwmarrin/discordgo)

## Features
* Copy embeds from the HordeUnited Discord to another channel
* Ping Roles when a specific buff was announced (optional)
* Delete old copied embeds when they are older than 36 hous (optional)

## Configs
If you want to fork it or run your own bot you need to configure the following things:
* two configs (json) to configure the horde united discord channel and one for all the guilds
* you need to enable developer mode in discord to copy the IDs of rolls, channels and your discordID
* server with golang installed
* your own discord bot

Unfortunately no way to configure it via a web-interface or something. Just a small bot to copy some messages over.
