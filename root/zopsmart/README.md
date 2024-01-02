Implement an API in a test driven format with following requirements:
API: POST /user/{name}/count

eg: 1st call to POST /user/sample-name/count
Response:
{
    name: "sample-name"
    count: 1
}

eg: 2nd call to POST /user/sample-name/count
Response:
{
    name: "sample-name"
    count: 2
}


Maintain separation of concern between the handler and the counter implementation