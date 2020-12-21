aws dynamodb put-item \
    --table-name Music \
    --item \
        '{"Artist": {"S": "No One You Know"}, 
        "SongTitle": {"S": "Call Me Today"},  
        "AlbumTitle": {"S": "Somewhat Famous"}}' \
    --return-consumed-capacity TOTAL \
    --endpoint-url http://localhost:8000