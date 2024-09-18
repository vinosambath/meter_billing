Thought process

**Goals**:
- The goal is read a given meter readings and store the values into table.
- Q: Should we push per day day or the data as it is available in the file (every 5, 15 or 30 minutes).
Refer Product discussions and decisions

**Understanding of the NMI12 format**:
1. 100 marks the start of one unit processing
2. Record Indicator 200 has 
   3. NMI information 
   4. Interval length - the value should be 5, 15 or 30 (indicating minutes)
3. Record Indicator 300 
   4. Interval Date 
   5. Interval Values - the number of values should be 1440/interval length
4. 900 marks the end of one unit processing

**Product discussions and decisions**:
This data will be useful for users in two ways
1. Display the data usage per hour as a graph [SP Utilities do this]
   2. Users can see usage per hour in SP Utilities helping them to decode which device is consuming a lot of electricity
2. Display the per day usage of the meter

#2: Display the per day usage of the meter
The application could sum all the usage of the meter for a day and store that value. This is more memory efficient.
The logic in this application handles this scenario. 

*#1: Display the data usage per hour as a graph*
This will definitely be useful for the user to see the information as a graph of per hour usage. This comes with storage
cost. the logic in this application could be modified easily to handle this scenario also along with #2

**Technical Details**:

**Coding structure and details**:

1. Meter Reading Service is considered a different service. The db layer sits above meter reading service and inside src/
2. Configuration to read configs, loggers are all initiated at the start of the service
3. Header Events are parsed and executed based on the NMIIndicator. This is made an interface 
and different NMIindicators will have different executors
   4. For instance, header_event_200_record_parser.go has the knowledge to parse NMI and intervalLength info
   5. header_event_300_record_executor holds the responsibility to store the meter reading in db


**Performance of the system**:
1. The db is shared based on the "nmi". The logic only have 2 logical shards, but this could be increased.
   1. Shard computation logic
      2. Hash the nmi to int and modulo by number of shards (which is 2)
   2. Assumptions:
      3. We will query the data per nmi
   4. Pros:
      5. Sharing the db loads (writes and reads) across different databases
   6. Cons:
      7. [Assuming its not needed] Joins on other fields not possible
      8. Querying by other fields require reverse index
9. The whole CSV file is NOT loaded into memory. Rather the CSV reader uses bufferIO to read line by line

**What could be improved**:
1. More logging to be added. Error logs were added extensively, more info logs could be added for better debugging
2. Handling failures. The system adds logs, instead it could try to push it to a retry queue
3. Storing the last processed record in the db, so when restarted the system processes it from the next record
3. Adding graceful shutdown
4. Add more tests. Test for the parser exist (header_event_record_parser_test.go), there is always scope to add more test 
and make the system more reliable 