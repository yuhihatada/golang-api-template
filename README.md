# golang API template

## About

This project is a template for golang API.

Package in use

- gorilla/mux
- postgres driver

## How to setup

1. install golang
1. install postgres
1. clone this repository
1. create database & tables
   - type command in docs/query/create
1. type command `go run main.go` in cmd/

## Usage

### get user's information

You can get user information.

#### path

GET: `/user/get/{user_id}`  
({user_id} is user identifier number)

#### request param

- none

#### request example

```
curl --location 'localhost:8000/user/get/1'
```

#### response example

`status code 200`

```json
{
  "id": 1,
  "name": "Yuhi",
  "mail": "yuhi@example.com"
}
```

### add user

You can add user.

#### path

POST: `/user/add`

#### request param

- name: string, require
  - You can set the user's name.
- mail: string, require
  - You can set the user's mail address.

#### request example

```
curl --location 'localhost:8000/user/add' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "name": "Yuhi",
    "mail": "yuhi@example.com"
}'
```

#### response example

`status code 200`

### edit user

You can edit user information.

#### path

PUT: `/user/edit`

#### request param

- id: number, require
  - Sets the identifier of the user to edit.
- name: string, require
  - Set the name after the change.
- mail: string, require
  - Set the address after the change.

#### request example

```
curl --location --request PUT 'localhost:8000/user/add' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "id": 1,
    "name": "Yuhi",
    "mail": "yuhi@example.com"
}'
```

#### response example

`status code 200`

### delete user

You can delete user.

#### path

DELETE: `/user/delete/{user_id}`
({user_id} is user identifier number)

#### request param

- none

#### request example

```
curl --location --request DELETE 'localhost:8000/user/delete/1'
```

#### response example

`status code 200`

### caution

To use the API, it is necessary to set an API key and other settings to prevent unspecified people from tapping the API.

## Note

Gorilla mux is archived. It will be migrated to gin in the future.
