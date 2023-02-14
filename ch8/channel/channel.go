package channel

// ch <- x   -- send instruction
// x = <-ch  -- receive instruction and set it value on x
// <-ch      -- just receive instruction, the result is ignored
// close(ch) -- close the channel, new messages will throw panic

// ch = make(chan int) -- unbuffered channel
// ch = make(chan int, 0) -- unbuffered channel
// ch = make(chan int, 3) -- buffered channel with 3 capacity
