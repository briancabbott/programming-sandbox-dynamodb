

aws dynamodb create-table \
    --table-name Music \
    --attribute-definitions \
        AttributeName=Artist,AttributeType=S \
        AttributeName=SongTitle,AttributeType=S \
    --key-schema \
        AttributeName=Artist,KeyType=HASH \
        AttributeName=SongTitle,KeyType=RANGE \
    --provisioned-throughput \
        ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --endpoint-url http://localhost:8000



aws dynamodb describe-table --table-name Music --endpoint-url http://localhost:8000


aws dynamodb get-item --consistent-read \
    --table-name Music \
    --key '{ "Artist": {"S": "Acme Band"}, "SongTitle": {"S": "Happy Day"}}' \
    --endpoint-url http://localhost:8000


aws dynamodb update-item \
    --table-name Music \
    --key '{ "Artist": {"S": "Acme Band"}, "SongTitle": {"S": "Happy Day"}}' \
    --update-expression "SET AlbumTitle = :newval" \
    --expression-attribute-values '{":newval":{"S":"Updated Album Title"}}' \
    --return-values ALL_NEW \
    --endpoint-url http://localhost:8000



aws dynamodb query \
    --table-name Music \
    --key-condition-expression "Artist = :name" \
    --expression-attribute-values '{":name":{"S":"Acme Band"}}'
    --endpoint-url http://localhost:8000


aws dynamodb update-table \
    --table-name Music \
    --attribute-definitions AttributeName=AlbumTitle,AttributeType=S \
    --global-secondary-index-updates '[{"Create":{"IndexName": "AlbumTitle-index", \
                                        "KeySchema":[{"AttributeName": "AlbumTitle", \ 
                                                      "KeyType":"HASH"}], \
                                       "ProvisionedThroughput": {"ReadCapacityUnits": 10, \ 
                                                                 "WriteCapacityUnits": 5}, \
                                       "Projection":{"ProjectionType":"ALL"}}}]'

aws dynamodb update-table \
--table-name Music \
--attribute-definitions AttributeName=AlbumTitle,AttributeType=S \
--global-secondary-index-updates \
                    "[{\"Create\":{\"IndexName\": \"AlbumTitle-index\",\"KeySchema\":[{\"AttributeName\":
                     \"AlbumTitle\",\"KeyType\":\"HASH\"}], \
                    \"ProvisionedThroughput\": {\"ReadCapacityUnits\": 10, \"WriteCapacityUnits\": 5
                    },\"Projection\":{\"ProjectionType\":\"ALL\"}}}]" \
 --endpoint-url http://localhost:8000



aws dynamodb query \
    --table-name Music \
    --index-name AlbumTitle-index \
    --key-condition-expression "AlbumTitle = :name" \
    --expression-attribute-values '{":name":{"S":"Somewhat Famous"}}' \
    --endpoint-url http://localhost:8000


aws dynamodb delete-table --table-name Music --endpoint-url http://localhost:8000
