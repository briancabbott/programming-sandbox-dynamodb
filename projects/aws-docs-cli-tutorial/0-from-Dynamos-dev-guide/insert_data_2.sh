aws dynamodb put-item \
    --table-name Music \
    --item '{"Artist": {"S": "Acme Band"}, 
             "SongTitle": {"S": "Happy Day"}, 
             "AlbumTitle": {"S": "Songs About Life"} }' \
    --return-consumed-capacity TOTAL \
    --endpoint-url http://localhost:8000