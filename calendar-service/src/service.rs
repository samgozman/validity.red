#![allow(dead_code, unused_variables)]

pub mod calendar {
    // 1. Read file with ID (from request data) as filename
    // 2. Decode file with AES-256
    // - optional: check if file is valid (if possible)
    pub fn read() {}

    // 1. Create .ics structure from data
    // 2. Write .ics to a buffer
    // 3. Encode buffer with AES-256
    pub fn create() {}

    // 1. Write encoded buffer to file with ID (from request data) as filename
    // 2. Replace previous calendar file (if exists)
    pub fn write() {}
}
