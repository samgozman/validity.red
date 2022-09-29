#![allow(dead_code)]

pub mod calendar {
    use crate::calendar::CalendarEntity;
    use chrono::{TimeZone, Utc};
    use ics::properties::{Description, DtEnd, DtStart, Summary};
    use ics::{escape_text, Event, ICalendar};

    // 1. Read file with ID (from request data) as filename
    // 2. Decode file with AES-256
    // - optional: check if file is valid (if possible)
    pub fn read() {}

    /// It takes a vector of `CalendarEntity`s and returns an `ICalendar` with the events from the
    /// `CalendarEntity`s
    ///
    /// Arguments:
    ///
    /// * `calendar_events`: A vector of CalendarEntity structs.
    ///
    /// Returns:
    ///
    /// A calendar object
    pub fn create(calendar_events: Vec<CalendarEntity>) -> ICalendar<'static> {
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

        return calendar;
    }

    // 1. Write encoded buffer to file with ID (from request data) as filename
    // 2. Replace previous calendar file (if exists)
    pub fn write() {}

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
