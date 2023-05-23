pub mod calendar {
    use crate::calendar::CalendarEntity;
    use crate::encryptor::{decrypt, encrypt};
    use chrono::{LocalResult, TimeZone};
    use chrono_tz::Tz;
    use ics::properties::{Description, DtEnd, DtStart, Summary};
    use ics::{escape_text, Event, ICalendar};
    use std::env;
    use std::error::Error;
    use std::fs::File;
    use std::io::{BufReader, Read, Write};
    use std::path::Path;
    use tonic::Status;

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
    pub fn read(file_name: &str, iv: &[u8; 12]) -> Result<String, Status> {
        const FILE_PATH: &str = "data/";
        let binding = FILE_PATH.to_owned() + file_name;
        let path = Path::new(binding.as_str());

        if !path.exists() {
            return Err(Status::not_found("Calendar file not found"));
        }

        let file = File::open(path).expect("Failed to open calendar file");
        let mut buf_reader = BufReader::new(file);

        let mut file_data: Vec<u8> = Vec::new();
        buf_reader
            .read_to_end(&mut file_data)
            .expect("Failed to read file");

        // TODO: Find a better way to get ENV variables in Rust
        let env_key = env::var("ENCRYPTION_KEY").expect("Expected ENCRYPTION_KEY ENV to be set");
        let mut encryption_key: [u8; 32] = Default::default();
        encryption_key.copy_from_slice(env_key.as_bytes());

        let decrypted = decrypt(file_data.as_slice(), &encryption_key, iv);
        match decrypted {
            Ok(data) => Ok(data),
            Err(e) => Err(Status::internal(e.to_string())),
        }
    }

    /// It takes a vector of `CalendarEntity`s and returns a string containing the iCalendar data
    ///
    /// Arguments:
    ///
    /// * `calendar_events`: A vector of CalendarEntity structs.
    /// * `tz`: Timezone identifier. For example, `Europe/Paris`.
    ///
    /// Returns:
    ///
    /// A string of the calendar.
    pub fn create(calendar_events: &Vec<CalendarEntity>, tz: Tz) -> String {
        let mut calendar = ICalendar::new(
            "2.0",
            "-//validity.extr.app//Document expiration calendar 1.0//EN",
        );

        // TODO: Add refresh params to calendar:
        // URL:http://my.calendar/url
        // REFRESH-INTERVAL;VALUE=DURATION:PT12H

        for calendar_event in calendar_events {
            let event = self::create_event(calendar_event, tz);
            calendar.add_event(event);
        }

        calendar.to_string()
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
        const FILE_PATH: &str = "data/";
        let path = FILE_PATH.to_owned() + file_name;
        let path = Path::new(&path);
        let parent_folder = path.parent().unwrap();

        if !parent_folder.exists() {
            std::fs::create_dir_all(parent_folder).unwrap();
        }

        // TODO: Find a better way to get ENV variables in Rust
        let env_key = env::var("ENCRYPTION_KEY").expect("Expected ENCRYPTION_KEY ENV to be set");
        let mut encryption_key: [u8; 32] = Default::default();
        encryption_key.copy_from_slice(env_key.as_bytes());

        let encrypted = encrypt(data, &encryption_key, iv).expect("Encryption failed");

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
    /// * `tz`: Timezone identifier. For example, `Europe/Paris`.
    ///
    /// Returns:
    ///
    /// A new event is being returned.
    fn create_event(calendar_event: &CalendarEntity, tz: Tz) -> Event<'static> {
        // TODO: Use ics standard timezone options instead of manual offset
        // Convert timestamp to DateTime
        let dt_start = calendar_event.notification_date.clone().unwrap();
        let dt_start = match tz.timestamp_opt(dt_start.seconds, dt_start.nanos as u32) {
            LocalResult::Single(dt) => dt,
            LocalResult::None => panic!("Invalid timestamp"),
            LocalResult::Ambiguous(_, _) => panic!("Ambiguous timestamp"),
        };

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
            validity.extr.app",
            calendar_event.document_title,
        ))));

        // ? CATEGORIES?

        event
    }

    #[cfg(test)]
    mod tests {
        use super::*;
        use crate::calendar::CalendarEntity;
        use prost_types::Timestamp;
        use serial_test::serial;
        use std::env;

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
            let tz: Tz = "Asia/Tbilisi".parse().unwrap();
            let event = create_event(&calendar_event, tz);
            let expected = "\
                BEGIN:VEVENT\r\n\
                UID:e533947d-6f40-4f4f-b614-ddf70534c576\r\n\
                DTSTAMP:20210107T101320Z\r\n\
                DTSTART:20210107T101320Z\r\n\
                DTEND:20210107T111320Z\r\n\
                SUMMARY:Document title\r\n\
                DESCRIPTION:Document title\\nvalidity.extr.app\r\n\
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
            let tz: Tz = "Asia/Tbilisi".parse().unwrap();
            let calendar = create(&calendar_events, tz);

            let expected = "\
                BEGIN:VCALENDAR\r\n\
                VERSION:2.0\r\n\
                PRODID:-//validity.extr.app//Document expiration calendar 1.0//EN\r\n\
                BEGIN:VEVENT\r\n\
                UID:e533947d-6f40-4f4f-b614-ddf70534c576\r\n\
                DTSTAMP:20210107T101320Z\r\n\
                DTSTART:20210107T101320Z\r\n\
                DTEND:20210107T111320Z\r\n\
                SUMMARY:Document title\r\n\
                DESCRIPTION:Document title\\nvalidity.extr.app\r\n\
                END:VEVENT\r\n\
                END:VCALENDAR\r\n\
            "
            .to_string();

            assert_eq!(calendar, expected, "unexpected output");
        }

        #[test]
        #[serial]
        fn test_write() {
            env::set_var("ENCRYPTION_KEY", "12345678901234567890123456789012");

            let iv = b"123456789012";
            let file_name = "tmp/test.ics";

            write("Test data string".to_string(), file_name, &iv).unwrap();

            let path = Path::new("data/tmp/test.ics");
            assert!(path.exists(), "File does not exist");

            // Clear test files
            std::fs::remove_dir_all(path.parent().unwrap()).unwrap();
        }

        #[test]
        #[serial]
        fn test_read() {
            env::set_var("ENCRYPTION_KEY", "12345678901234567890123456789012");

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
