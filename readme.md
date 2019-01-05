# gotube

a music server for linux-based systems that queues up and plays songs from youtube. inspired by [mps-youtube](https://github.com/mps-youtube/mps-youtube)

## client

to actually listen to anything you'll need a client. currently the only available one is [here](https://github.com/monodokimes/gotube-client) but i'm planning on making a cli version too.

## requirements

* [golang](https://golang.org/), until prebuilt releases are a thing
* [youtube-dl](https://rg3.github.io/youtube-dl)

## searching

by default gotube uses `youtube-dl`'s simulation options to search for tracks. this is very slow, and increases linearly with the number of search results returned. as such it is possible to search using youtube's api directly, but an api key must be provided. [go here to get one](https://developers.google.com/youtube/v3/getting-started).
