fn main() {
    // 1. Start gRPC server
    // 2. Listen for WriteCalendar and ReadCalendar requests

    // WriteCalendar handler
    // 1. Recieve combined calendar data from broker
    // 2. Create .ics structure from data
    // 3. Write .ics to a buffer
    // 4. Encode buffer with AES-256
    // 5. Write encoded buffer to file with ID (from request data) as filename
    // 7. Replace previous calendar file (if exists)

    // ReadCalendar handler
    // 1. Read file with ID (from request data) as filename
    // 2. Decode file with AES-256
    // - optional: check if file is valid (if possible)
    // 3. Send decoded structure as a gRPC stream to the broker
}
