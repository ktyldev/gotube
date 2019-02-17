server="localhost:6969"
contenttype="content-type:application/json"

declare -a data_array=(
  # add to the back of the queue (totorro - home alone)
  '{"id":"xFWVFu2ASbE","index":-1}'
  # add to the front of the queue (clever girl - no drum and bass in the jazz room")
  '{"id":"9ANzhPJZ2U0","index":0}'
  # add something in the middle (Black Hill & Silent Island - Tales of the night forest [Full Album])
  '{"id":"mTLunRuCGQQ","index":1}'
  # add something to the back
  '{"id":"xFWVFu2ASbE","index":-1}'
)

echo_queue() {
  echo ""
  echo "QUEUE"
  echo ""
  curl "$server/queue"
}

echo "adding songs to queue..."
for d in "${data_array[@]}"
do
  curl -H $contenttype -d $d "$server/add"
done

echo_queue
echo ""
echo "MOVING INDEX 1 TO INDEX 2"
echo ""
curl -H $contenttype -d '{"oldIndex":1, "newIndex":2}' "$server/move"

echo_queue
