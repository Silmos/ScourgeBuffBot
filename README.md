# ScourgeBuffBot
## Discord Bot

![alt text](https://i.imgur.com/4Ge69KX.png)

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

## Contact
If you are also on Transcendence-Horde and want updates from my bot, just contact me via Discord (Silmos#3481). I need the following data:
* your guildname
* the channelID for the copied embed messages from HordeUnited
* your guildID
* Should the bot mention special roles? (True or False)
* RoleID for rend buffs (for pinging roles if rend buffs are known)
* RoleID for hakkar buffs (for pinging roles if hakkar buffs are known)
* RoleID for dragon buffs NEF/ONY (for pinging if one of the buffs are known)

I recommend you to setup your own carl-bot to create a reaction role embed. Create three emotes for hakkar, rend and dragon buffs. If a user click them, assign the role to them so they get directly pinged if a buff is known.

You also need to invite the bot to your server.

https://discord.com/oauth2/authorize?client_id=814948099638296657&scope=bot&permissions=0

The bot will not have any permissions, so you need to enable the following permissions for your announcement channel:
* Show Channel
* Send Messages
* Mention all roles (if you want the bot to mention roles)
* Manage Messages (if you want to delete old messages from the bot automatically)
* Show Chat log (don't really think this is needed, but does not hurt)

To see the IDs of channels, guilds, messages, rolls etc. you need to enable developer mode for discord.
Official Documentation: https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-

## Credits
Programmed (badly) by myself

discordgo - https://github.com/bwmarrin/discordgo
