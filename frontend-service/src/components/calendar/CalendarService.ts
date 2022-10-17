import { QueryMaker, type IResponse } from "@/services/QueryMaker";
import type { ICalendarNotification } from "./interfaces/ICalendarNotification";

interface ICalendarResponse extends IResponse {
  data: {
    calendar: ICalendarNotification[];
  };
}

export class CalendarService {
  // TODO: Add params for pagination by month
  /**
   * Get array of calendar notifications for current user
   * @returns array of calendar notifications
   */
  public static async getCalendar(): Promise<ICalendarNotification[]> {
    const res = await new QueryMaker({
      route: "/calendar",
    }).get<ICalendarResponse>();
    const { error, message, data } = res.data;

    if (error) {
      throw new Error(message);
    }

    return data.calendar;
  }

  /**
   * Create calendar for specified Date (today) from events array
   * @param today - Date object
   * @param events - array of calendar notifications
   * @returns map of events with day number as key
   */
  public static createCalendar(
    today: Date,
    events: ICalendarNotification[]
  ): Map<number, ICalendarNotification[]> {
    const daysInMonth = new Date(
      today.getFullYear(),
      today.getMonth() + 1,
      0
    ).getDate();

    const calendarMonth = new Map<number, ICalendarNotification[]>();
    for (let day = 1; day < daysInMonth; day++) {
      // Find associated notifications for this day
      const notifications = events.filter((n) => {
        const notificationDate = new Date(n.notificationDate.split("T")[0]);
        const calendarDay = new Date(
          today.getFullYear(),
          today.getMonth(),
          day
        );

        return (
          notificationDate.toLocaleDateString("en-GB") ===
          calendarDay.toLocaleDateString("en-GB")
        );
      });

      calendarMonth.set(day, notifications);
    }

    return calendarMonth;
  }
}
