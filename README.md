# Hololive Livestream Notification Bot

Sends notifications when livestreams from hololive are about to begin. (15 minutes prior and on time)

Requires a server text room with the name `gobot` and `gobot-live`

Bot link: <a>TBD</a>

## Using the project?

Requirements: MongoDB database (e.g. Free MongoDB Atlas cluster), Heroku, golang

### Host on Heroku

I hosted my bot on Heroku. Here's a dev guide on setting it up: https://devcenter.heroku.com/articles/git

### Set up environment variables / config vars
PORT=port#

COMMAND_PREFIX=command prefix

DISCORD_BOT_TOKEN=really long discord bot token

DEVELOPER_IDS=your discord id so you can get DMs from the bot about processing feeds

DEVELOPER_MODE=ON (enables the DMs that help debug issues)

WEBPAGE=the heroku webpage (e.g. your-gobot-app.herokuapp.com)

MONGO_CONNECTION_URL=connection url to a mongodb instance

DEVELOPER_MODE=ON

### Set up MongoDB Database and Collections

You can get a free cluster from MongoDB Atlas.
You can use MongoDB Compass to interface with the cluster.
You will need to create a hololive-en database. Contained inside are three collections `feeds`, `scheduledLivestreams`, `subscriptions`. You will need to create all three.
The only documents you will need to manually generate is in the `feeds` collection. Each document should describe the vtuber being tracked.

Example:
```
{
    "firstName": "gawr",
    "lastName": "gura",
    "topicURL": "https://www.youtube.com/xml/feeds/videos.xml?channel_id=UCoSrY_IQQVpmIRZ9Xf-y93g",
    "group": "en",
    "generation": 1
}
```

### Questions?

Discord: <b>きつね#1040</b>

Twitter: <b>@SheavinNou</b>
