import { QueryMaker, type IResponse } from "@/services/QueryMaker";
import type { IDashboardStats } from "./interfaces/IDashboardStats";

interface IGetStatsResponse extends IResponse {
  data: IDashboardStats;
}

export class DashboardService {
  public static async getStats(): Promise<IDashboardStats> {
    const res = await new QueryMaker({
      route: "/documents/statistics",
    }).get<IGetStatsResponse>();
    const { error, message, data } = res.data;

    if (error) {
      throw new Error(message);
    }

    for (const d of data.latestDocuments) {
      // Return date in YYYY-MM-DD format
      d.expiresAt = new Date(d.expiresAt).toISOString().substring(0, 10);
    }

    // Sort used types in descending order
    data.usedTypes = data.usedTypes.sort((a, b) =>
      a.count < b.count ? 1 : -1
    );

    return data;
  }

  public static async getIcsFile(calendarId: string): Promise<string> {
    const res = await new QueryMaker({
      route: `/ics/${calendarId}`,
    }).get<any>();

    return res.data;
  }
}
