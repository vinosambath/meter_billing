Thought process

Goals:
The goal is read a given meter readings and store the values into table.

Understanding of the NMI 12 format:
1. 100 marks the start of one unit processing
2. Record Indicator 200 has 
   a. NMI information
   b. Interval length - the value should be 5, 15 or 30 (indicating minutes)
3. Record Indicator 300
   a. Interval Date
   b. Interval Values - the number of values should be 1440/interval length
4. 900 marks the end of one unit processing

Product discussions and decisions:
This data will be useful for users in two ways
1. Display the data usage per hour as a graph
    a. As a user I can see usage per hour in SP Utilities helping me to decode which device is consuming a lot of electricity
2. Display the per day usage of the user

I assume #2 is the requirement given and #1 is a good to build feature for the user.

Technical Details:


Coding structure and details:

1. Graceful Shutdown
2. 


Performance of the system
1. The db is shared based on the "nmi". The logic only have 2 logical shards, but this could be increased.
   1. Shard computation logic
      2. Hash the nmi to int and modulo by number of shards (which is 2)
   2. Assumptions:
      3. We will query the data per nmi
   4. Pros:
      5. Sharing the db load (writes and reads) across different databases
   6. Cons:
      7. Joins on other fields not possible
      8. Querying by other fields require reverse index
9. Reading CSV files in memory
