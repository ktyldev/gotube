/ping
/version

/queue/add
    POST {id}
    return status

/queue/next->/queue/advance
    POST

/queue/clear
    POST

/queue/remove/{index}
    POST

[deprecated]/queue/top

/queue
    GET
    [{id, title}]

/stream/{id}
/search

/info/{id}
    title
    artwork

save songs as [id].webm
