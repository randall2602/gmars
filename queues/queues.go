// package implements the Task Queues of the Memory Array Redcode Simulator (MARS)
package queues

import "fmt"

var QUEUES_DESCRIPTION = 'the MARS has a process queue, a list of processes that are executed
 repeatedly in the order in which they were started. New processes created by SPL are added
 just after the current process, while those that execute a DAT will be removed from the queue.
 If all the processes die, the warrior will lose.'

func WhatIsQueues() string {
  return TASKS_DESCRIPTION
}
