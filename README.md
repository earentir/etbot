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

- bofh
  - Pulls a BOFH quote from an API
- crypto \*
  - Does crypto currency conversions
- etbdown
  - Does a bot reload (not fully yet)
- exchange \*
  - Does FIAT currency conversions
- fr $
  - Adds Issues in git hub
- joke
  - You can add/search/delete jokes
- level
  - Outputs the users level/type
- love
  - Users can add something they love and it will be appended in the !so
- lurk
  - Users can lurk, it also allows to type a reason, the user gets unlurked the
    next time they type something
- project
  - Current working project (future sync with stream name/description as an
    option)
- quote
  - Add/Delete/Search Quotes
- savesettings
  - Forces a flush of all in memory settings to files
- so \*
  - Will shout out a user, if they have an account all their socials will be
    printed out including what they love (fully customizable message soon)
- socials
  - Prints streamers socials
- time -Prints the streamers time or the time of any country if you append the
  tz city
- tmdb \*
  - Pulls data for tv/movie from tmdb
- updsoc
  - Allows users to add socials to their accounts
- user
  - Add/Delete accounts from the system (will sync twitch type soon)
- version
  - Prints the bots version (build date for now)
- weather \*
  - Prints the weather in the streamer city or for any city if appended to the
    command
- year
  - Prints some stats about the current year (day, % of the year left etc)
- yoke
  - Prints one of my favorite jokes, its an API call
- zoe
  - Pet feeding and petting mechanism, allows users to add treats/petting time
    to a que

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
