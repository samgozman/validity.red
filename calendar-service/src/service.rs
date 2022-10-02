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
    ///
    /// Returns:
    ///
    /// A String containing the contents of the file or an ([`Err`]).
    pub fn read(file_name: &str) -> Result<String, Box<dyn Error>> {
        // TODO: Read from env
        const FILE_PATH: &str = "data/";
        let path = FILE_PATH.to_owned() + file_name;

        let file = File::open(path).expect("File not found");
        let mut buf_reader = BufReader::new(file);

        let mut file_data: Vec<u8> = Vec::new();
        buf_reader
            .read_to_end(&mut file_data)
            .expect("Failed to read file");

        // TODO: Read from Proto
        let iv: &[u8; 12] = b"123456789012";
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
    pub fn create(calendar_events: Vec<CalendarEntity>) -> String {
        // TODO: recieve calendar_events as a pointer
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
    ///
    /// Returns:
    ///
    /// A Result that either success ([`Ok`]) or failure ([`Err`])
    pub fn write(data: String, file_name: &str) -> Result<(), Box<dyn Error>> {
        // TODO: Read from env
        const FILE_PATH: &str = "data/";
        let path = FILE_PATH.to_owned() + file_name;
        let path = std::path::Path::new(&path);
        let parent_folder = path.parent().unwrap();

        if !parent_folder.exists() {
            std::fs::create_dir_all(parent_folder).unwrap();
        }

        // TODO: Read from Proto
        let iv: &[u8; 12] = b"123456789012";

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
    fn create_event(calendar_event: CalendarEntity) -> Event<'static> {
        // Convert timestamp to DateTime
        let dt_start = calendar_event.notification_date.unwrap();
        let dt_start = Utc.timestamp(dt_start.seconds, dt_start.nanos as u32);
        // End date is 1 hour after start date
        let dt_end = dt_start + chrono::Duration::hours(1);

        // Convert date to ICS ISO 8601 format
        let dt_start = dt_start.format("%Y%m%dT%H%M%SZ").to_string();
        let dt_end = dt_end.format("%Y%m%dT%H%M%SZ").to_string();

        let mut event = Event::new(calendar_event.notification_id, dt_start.clone());
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
}
