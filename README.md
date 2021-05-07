### Evermos

1. Describe what you think happened that caused those bad reviews during our 12.12 event and why it happened.

```text
I think the bad review happened because customers who are disappointed because they cannot buy their products. They already 
succeed to check out their order but it cant processed because of miss stock in inventory. Negative value in inventory 
quantity usually occurred because of race condition.
```

2. Based on your analysis, propose a solution that will prevent the incidents from occurring again.

```text
Race condition can prevented with lock on row of table. We will lock inv qty so when someone access the row it will wait for
update/process to happened first. If update lock release it will have accurate inventory quantity
```

3. Based on your proposed solution, build a Proof of Concept that demonstrates technically how your solution will work

![alt text](https://i.imgur.com/GiPVU87.png "System Architecture")

### Architecture

In this repo we'll use Gateway to handle all http request and pass the request to services (inventory)

### Folder structure

services will go to their separate folder domain, example: gateway; inventory_service.

pkg will used as library helper to easy access to all repo

### Prerequisite

if you already have mysql server you can skip to number 2

1. you need to start mysql server via docker. start it with

```bash
make start
```

2. dump sql db script in folder ./setup/evermos_20210507_232026.sql. if database setup ready we can move to testing

### Testing

In this repo there are 2 available test.

1. Test reduce qty inventory with race condition
2. Test reduce qty inventory without race condition

Testing method:

- Test will spawn 50 go routine to simulate 50 person buy at same product simultaneously
- product id will only available with 5 qty
- the expected result is product qty will be 0 with no negative value

Running test To run the test type command below

1. With no race condition

```bash
go test ./inventory_service/repo -tags integration -count 1 -v -run=^TestInventoryRepo_UpdateInventoryQty
```

2. With race condition

```bash
go test ./inventory_service/repo -tags integration -count 1 -v -run=^TestInventoryRepo_UpdateInventoryQtyRaceCondition
```

### API

To see available the api please export postman at ./setup/postman.json

Run the repo:

terminal 1 (gateway)

```bash
cd gateway && go run main.go
```

terminal 2: inventory service

```bash
 cd inventory_service && go run main.go
```

HTTP can be accessed at localhost:3004 and

gRPC can be accessed at localhost:3001 