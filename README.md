# etbot aka Ear's Twitch Bot

## Intro

The bot will provide general bot services for Twitch, twitch integration (read
only for now for security reasons, will update with write when more stable),
discord integration (in the future) and OBS integration (in the future) through
a browser service. The plan is to have a single binary running all services and
multiplexing from there.

### Permission System

Each registered user has a user type, each user type has a power level and a
time multiplier, each command has a minimum power level attached and a timeout.

### System Commands

System commands cannot be removed, they can be dissabled though and most have
the ability to customize messages (all wii have that in the future)

- joke
- bofh
- crypto \*
- etbdown
- exchange \*
- fr $
- level
- love
- lurk
- project
- quote
- savesettings
- so \*
- socials
- time
- tmdb \*
- updsoc
- user
- version
- weather \*
- year
- yoke
- zoe

Commands with \* require an API key from a third party. Commands with a $
require a local client.

### User Commands

User commands are of 4 types punchline, varpunchline, counter and tree(not
implemented yet), the default are there just as an example(and I use them in my
channel)

- ban
- f
- hi
- hype
- mic
- nvidia
- oil
- putin
- sudo
- unban
- yogurt
