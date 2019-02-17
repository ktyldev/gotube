server="localhost:6969"
contenttype="content-type:application/json"

declare -a data_array=(
  # add to the back of the queue (totorro - home alone)
  '{"id":"xFWVFu2ASbE","index":-1}'
  # add to the front of the queue (clever girl - no drum and bass in the jazz room")
  '{"id":"9ANzhPJZ2U0","index":0}'
  # add something in the middle (Black Hill & Silent Island - Tales of the night forest [Full Album])
  '{"id":"mTLunRuCGQQ","index":1}'
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
  curl -H $contenttype -d $d "$server/queue/add"
done

echo_queue
echo ""
echo "REMOVING INDEX 1"
echo ""
curl -H $contenttype -d '{"index":1}' "$server/queue/remove"

echo_queue
