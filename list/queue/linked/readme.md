# linked-list based queue

Linked-list based queues must be doubly-linked lists, with previous and next references, 
because we need to be able to read from the front of the queue and delete it, while inserting at the end. 
It could be implemented as a singly-linked list, but it would require sacrificing performance on one or the other operations.
