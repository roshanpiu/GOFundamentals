//maps are reference types 
//maps are passed by reference
//maps are not safe for concurrency cus we dont want to write and read to maps at the same time
//maps are cheap to pass around
//specify size for large maps for better performance ->> make(map[<keytype>]<valueType>,size)