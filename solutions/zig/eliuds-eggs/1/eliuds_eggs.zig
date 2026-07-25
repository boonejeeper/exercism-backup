pub fn eggCount(number: usize) usize {
    var counter: usize = 0;
    var workingNumber: usize = number;
    while (workingNumber > 0) {
        if (workingNumber & 1 == 1) {
            counter += 1;
        }
        workingNumber = workingNumber >> 1;
    }
    return counter;
}
