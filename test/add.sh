server="localhost:6969"

declare -a data_array=(
  # add to the back of the queue (totorro - home alone)
  '{"id":"xFWVFu2ASbE","index":-1}'
  # add to the front of the queue (clever girl - no drum and bass in the jazz room")
  '{"id":"9ANzhPJZ2U0","index":0}'
  # add some bulk...
  '{"id":"xFWVFu2ASbE","index":-1}'
  '{"id":"xFWVFu2ASbE","index":-1}'
  # add something in the middle (Black Hill & Silent Island - Tales of the night forest [Full Album])
  '{"id":"mTLunRuCGQQ","index":1}'
)

for d in "${data_array[@]}"
do
  echo ""
  echo "POST $d"
  echo ""
  curl -H "content-type:application/json" -d $d "$server/queue/add"
  echo ""
  echo "RESPONSE"
  echo ""
  curl "$server/queue"
  echo ""
done



