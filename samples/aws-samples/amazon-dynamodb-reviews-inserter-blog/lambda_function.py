import json
import boto3
import os
import datetime



def lambda_handler(event, context):
    
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table(os.environ['table_name'])
    
   
    with table.batch_writer() as batch:
        for item in event:
            if "/" in item["review_date"]:
                item["review_date"] = datetime.datetime.strptime(item["review_date"], '%m/%d/%y').strftime('%Y-%m-%d')
            item["review_date-review_id"] = item["review_date"] + item["review_id"]
            item["customer_id"] = str(item["customer_id"])

            batch.put_item(item)
    
    return {
        'statusCode': 200,
        'body': json.dumps('Reviews inserted!')
    }
