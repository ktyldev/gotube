# gotube

a music server that queues up and plays songs from youtube. inspired by [mps-youtube](https://github.com/mps-youtube/mps-youtube)

## requirements

* [youtube-dl](https://rg3.github.io/youtube-dl)

## searching

by default gotube uses `youtube-dl`'s simulation options to search for tracks. this is very slow, and increases linearly with the number of search results returned. as such it is possible to search using youtube's api directly, but an api key must be provided. [go here to get one](https://developers.google.com/youtube/v3/getting-started).
