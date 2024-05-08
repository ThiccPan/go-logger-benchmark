STARTTS=$EPOCHSECONDS
if [[ $1 == "t-getall" ]]
then
    k6 run ./get-items.js --out csv=test-result/test-getall-$2.csv --summary-export=test-result/summary-getall-$2.json
elif ([ $1 == "t-edit" ])
then
    k6 run ./edit-item.js --out csv=test-result/test-edit-$2.csv --summary-export=test-result/summary-edit-$2.json
elif ([ $1 == "t-add" ])
then
    k6 run ./add-item.js --out csv=test-result/test-add-$2.csv --summary-export=test-result/summary-add-$2.json
else
    echo "command not registered"
    exit 1
fi
ENDTS=$EPOCHSECONDS
curl http://localhost:19999/api/v1/data\?chart\=app.logging-experiment_cpu_utilization\&after\=$STARTTS\&before\=$ENDTS > test-result/cpu-$1-$2.json
curl http://localhost:19999/api/v1/data\?chart\=app.logging-experiment_mem_private_usage\&after\=$STARTTS\&before\=$ENDTS > test-result/mem-$1-$2.json