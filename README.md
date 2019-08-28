Introduction

This is a blockchain project which pulls a certain number of blocks from a public blockchain network, stores the data into a database and fetches the transaction details given the block number.

Steps to follow:
1. Download the zip or clone the repository into your local system.
2. Go to the folder.
3. Database used is mysql, you have to provide the username, password, hostname, port number and the schema. Design pattern DAO is used to insert and fetch the data to/from database.
4. Schema Name: blockchain, table name: transaction_structs
    column names: from_addr, to_addr, block_number, transaction_number.
5. To retrieve the user transactions, http server is used which would run on localhost:8081/{blocknumber}
6. all the transactions would be displayed onto the browser screen which are associated with a particular block, if there are no such transactions, an error message would be displayed.
7. To execute the program, you can run command "go run main.go" or you can build the binary using command "go build main.go" and the execute the binary by executing "./main".