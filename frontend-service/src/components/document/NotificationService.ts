import { QueryMaker, type IResponse } from "@/services/QueryMaker";
import type { INotification } from "./interfaces/INotification";

interface NotificationDeletePayload {
  id: string;
  documentId: string;
}

interface NotificationGetAllResponse extends IResponse {
  data: {
    notifications: INotification[];
  };
}

interface NotificationPayload {
  date: Date;
  documentId: string;
}

export class NotificationService {
  /**
   * Delete notification object
   * @param params
   */
  public static async deleteOne(
    params: NotificationDeletePayload
  ): Promise<void> {
    const res = await new QueryMaker({
      route: `/documents/${params.documentId}/notifications/delete/${params.id}`,
    }).delete();
    const { error, message } = res.data;

    if (error) {
      throw new Error(message);
    }
  }

  /**
   * Delete notification object
   * @returns
   */
  public static async getAll(documentId: string): Promise<INotification[]> {
    const res = await new QueryMaker({
      route: `/documents/${documentId}/notifications`,
    }).get<NotificationGetAllResponse>();
    const { error, message, data } = res.data;

    if (error) {
      throw new Error(message);
    }

    return data.notifications || [];
  }

  /**
   * Create notification object
   * @param notification
   */
  public static async createOne(
    notification: NotificationPayload
  ): Promise<void> {
    const payload = JSON.stringify({ date: notification.date });

    const res = await new QueryMaker({
      route: `/documents/${notification.documentId}/notifications/create`,
      payload,
    }).post<NotificationGetAllResponse>();
    const { error, message } = res.data;

    if (error) {
      throw new Error(message);
    }
  }
}
