import { QueryMaker, type IResponse } from "@/services/QueryMaker";
import type { IDashboardStats } from "./interfaces/IDashboardStats";

interface IGetStatsResponse extends IDashboardStats, IResponse {}

export class DashboardService {
  public static async getStats(): Promise<IDashboardStats> {
    const res = await new QueryMaker({
      route: "/documents/statistics",
    }).get<IGetStatsResponse>();
    const {
      error,
      message,
      totalDocuments,
      totalNotifications,
      usedTypes,
      latestDocuments,
    } = res.data;

    if (error) {
      throw new Error(message);
    }

    for (const d of latestDocuments) {
      // Return date in YYYY-MM-DD format
      d.expiresAt = new Date(d.expiresAt).toISOString().substring(0, 10);
    }

    return {
      totalDocuments,
      totalNotifications,
      // Sort used types in descending order
      usedTypes: usedTypes.sort((a, b) => (a.count < b.count ? 1 : -1)),
      latestDocuments,
    };
  }

  public static async getIcsFile(calendarId: string): Promise<string> {
    const res = await new QueryMaker({
      route: `/ics/${calendarId}`,
    }).get<string>();

    return res.data;
  }
}
