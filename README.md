# enterpret
customer feedback ingestion
```
curl --location 'http://localhost:8080/pull?id=1&startDate=2024-07-18&endDate=2024-08-18' => To pull feedbacks

curl --location 'http://localhost:8080/push' \
--header 'Content-Type: application/json' \
--data '{
    "id": "12",
    "sourceId": "12"
}' => To push feedback

curl --location 'http://localhost:8080/findAll' => To fetch all feedbacks

```



![feedback-ingestion](https://github.com/user-attachments/assets/b49366ff-c7a4-45dc-8dda-49238447cfd5)
