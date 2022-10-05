pub mod calendar {
    use crate::calendar::CalendarEntity;
    use crate::encryptor::{decrypt, encrypt};
    use chrono::{TimeZone, Utc};
    use ics::properties::{Description, DtEnd, DtStart, Summary};
    use ics::{escape_text, Event, ICalendar};
    use std::error::Error;
    use std::fs::File;
    use std::io::{BufReader, Read, Write};

    // TODO: Read from env
    const KEY: &[u8; 32] = b"A%D*G-KaPdSgVkYp3s6v9y$B?E(H+MbQ";

    /// Read a file and return the contents as a string
    ///
    /// Arguments:
    ///
    /// * `file_name`: The name of the file to read.
    /// * `iv`: Initialization vector. This is a random value that is used to ensure that the same plaintext
    ///
    /// Returns:
    ///
    /// A String containing the contents of the file or an ([`Err`]).
    pub fn read(file_name: &str, iv: &[u8; 12]) -> Result<String, Box<dyn Error>> {
        // TODO: Read from env
        const FILE_PATH: &str = "data/";
        let path = FILE_PATH.to_owned() + file_name;

        let file = File::open(path).expect("File not found");
        let mut buf_reader = BufReader::new(file);

        let mut file_data: Vec<u8> = Vec::new();
        buf_reader
            .read_to_end(&mut file_data)
            .expect("Failed to read file");

        let decrypted = decrypt(file_data.as_slice(), KEY, iv);

        Ok(decrypted)
    }

    /// It takes a vector of `CalendarEntity`s and returns a string containing the iCalendar data
    ///
    /// Arguments:
    ///
    /// * `calendar_events`: A vector of CalendarEntity structs.
    ///
    /// Returns:
    ///
    /// A string of the calendar.
    pub fn create(calendar_events: &Vec<CalendarEntity>) -> String {
        let mut calendar = ICalendar::new(
            "2.0",
            "-//Validity.Red//Document expiration calendar 1.0//EN",
        );

        // TODO: Add refresh params to calendar:
        // URL:http://my.calendar/url
        // REFRESH-INTERVAL;VALUE=DURATION:PT12H

        for calendar_event in calendar_events {
            let event = self::create_event(calendar_event);
            calendar.add_event(event);
        }

        return calendar.to_string();
    }

    /// It creates a file with the given name and writes the given data to it
    ///
    /// Arguments:
    ///
    /// * `data`: String - This is the data that we want to write to the file.
    /// * `file_name`: The name of the file to write to.
    /// * `iv`: Initialization vector. This is a random value that is used to ensure that the same plaintext
    ///
    /// Returns:
    ///
    /// A Result that either success ([`Ok`]) or failure ([`Err`])
    pub fn write(data: String, file_name: &str, iv: &[u8; 12]) -> Result<(), Box<dyn Error>> {
        // TODO: Read from env
        const FILE_PATH: &str = "data/";
        let path = FILE_PATH.to_owned() + file_name;
        let path = std::path::Path::new(&path);
        let parent_folder = path.parent().unwrap();

        if !parent_folder.exists() {
            std::fs::create_dir_all(parent_folder).unwrap();
        }

        let encrypted = encrypt(data, KEY, iv).expect("Encryption failed");

        let mut file = File::create(path).expect("Unable to create file");
        file.write_all(encrypted.as_slice())
            .expect("Unable to write data");

        Ok(())
    }

    /// It takes a `CalendarEntity` and returns an iCal `Event` with the same data
    ///
    /// Arguments:
    ///
    /// * `calendar_event`: The CalendarEntity struct that contains the data for the event.
    ///
    /// Returns:
    ///
    /// A new event is being returned.
    fn create_event(calendar_event: &CalendarEntity) -> Event<'static> {
        // Convert timestamp to DateTime
        let dt_start = calendar_event.notification_date.clone().unwrap();
        let dt_start = Utc.timestamp(dt_start.seconds, dt_start.nanos as u32);
        // End date is 1 hour after start date
        let dt_end = dt_start + chrono::Duration::hours(1);

        // Convert date to ICS ISO 8601 format
        let dt_start = dt_start.format("%Y%m%dT%H%M%SZ").to_string();
        let dt_end = dt_end.format("%Y%m%dT%H%M%SZ").to_string();

        let mut event = Event::new(calendar_event.notification_id.clone(), dt_start.clone());
        event.push(DtStart::new(dt_start));
        event.push(DtEnd::new(dt_end));
        event.push(Summary::new(calendar_event.document_title.clone()));
        event.push(Description::new(escape_text(format!(
            "{}\n\
            Validity.Red",
            calendar_event.document_title,
        ))));

        return event;
    }

    #[cfg(test)]
    mod tests {
        use super::*;
        use crate::calendar::CalendarEntity;
        use prost_types::Timestamp;

        #[test]
        fn test_create_event() {
            let calendar_event = CalendarEntity {
                notification_id: "e533947d-6f40-4f4f-b614-ddf70534c576".to_string(),
                document_title: "Document title".to_string(),
                notification_date: Some(Timestamp {
                    seconds: 1610000000,
                    nanos: 0,
                }),
                document_id: "8d3c3c83-b4e1-466e-ad71-29948479b38e".to_string(),
                expires_at: Some(Timestamp {
                    seconds: 1630000000,
                    nanos: 0,
                }),
            };

            let event = create_event(&calendar_event);
            let expected = "\
                BEGIN:VEVENT\r\n\
                UID:e533947d-6f40-4f4f-b614-ddf70534c576\r\n\
                DTSTAMP:20210107T061320Z\r\n\
                DTSTART:20210107T061320Z\r\n\
                DTEND:20210107T071320Z\r\n\
                SUMMARY:Document title\r\n\
                DESCRIPTION:Document title\\nValidity.Red\r\n\
                END:VEVENT\r\n\
            "
            .to_string();

            assert_eq!(event.to_string(), expected, "unexpected output");
        }

        #[test]
        fn test_create() {
            let calendar_events = vec![CalendarEntity {
                notification_id: "e533947d-6f40-4f4f-b614-ddf70534c576".to_string(),
                document_title: "Document title".to_string(),
                notification_date: Some(Timestamp {
                    seconds: 1610000000,
                    nanos: 0,
                }),
                document_id: "8d3c3c83-b4e1-466e-ad71-29948479b38e".to_string(),
                expires_at: Some(Timestamp {
                    seconds: 1630000000,
                    nanos: 0,
                }),
            }];

            let calendar = create(&calendar_events);

            let expected = "\
                BEGIN:VCALENDAR\r\n\
                VERSION:2.0\r\n\
                PRODID:-//Validity.Red//Document expiration calendar 1.0//EN\r\n\
                BEGIN:VEVENT\r\n\
                UID:e533947d-6f40-4f4f-b614-ddf70534c576\r\n\
                DTSTAMP:20210107T061320Z\r\n\
                DTSTART:20210107T061320Z\r\n\
                DTEND:20210107T071320Z\r\n\
                SUMMARY:Document title\r\n\
                DESCRIPTION:Document title\\nValidity.Red\r\n\
                END:VEVENT\r\n\
                END:VCALENDAR\r\n\
            "
            .to_string();

            assert_eq!(calendar, expected, "unexpected output");
        }

        #[test]
        fn test_write() {
            let iv = b"123456789012";
            let file_name = "tmp/test.ics";

            write("Test data string".to_string(), file_name, &iv).unwrap();

            let path = std::path::Path::new("data/tmp/test.ics");
            assert!(path.exists(), "File does not exist");

            // Clear test files
            std::fs::remove_dir_all(path.parent().unwrap()).unwrap();
        }

        #[test]
        fn test_read() {
            let iv = b"123456789012";
            let file_name = "tmp/test.ics";

            write("Test data string".to_string(), file_name, &iv).unwrap();

            let data = read(file_name, &iv).unwrap();
            assert_eq!(data, "Test data string", "Unexpected data");

            // Clear test files
            std::fs::remove_dir_all("data/tmp").unwrap();
        }
    }
}
