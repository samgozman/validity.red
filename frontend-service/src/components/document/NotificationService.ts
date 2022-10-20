import { QueryMaker, type IResponse } from "@/services/QueryMaker";
import type { INotification } from "./interfaces/INotification";

interface INotificationDeletePayload {
  id: string;
  documentId: string;
}

interface INotificationGetAllResponse extends IResponse {
  notifications: INotification[];
}

interface INotificationPayload {
  date: Date;
  documentId: string;
}

export class NotificationService {
  /**
   * Delete notification object
   * @param params
   */
  public static async deleteOne(
    params: INotificationDeletePayload
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
    }).get<INotificationGetAllResponse>();
    const { error, message, notifications } = res.data;

    if (error) {
      throw new Error(message);
    }

    return notifications || [];
  }

  /**
   * Create notification object
   * @param notification
   */
  public static async createOne(
    notification: INotificationPayload
  ): Promise<void> {
    const payload = JSON.stringify({ date: notification.date });

    const res = await new QueryMaker({
      route: `/documents/${notification.documentId}/notifications/create`,
      payload,
    }).post<INotificationGetAllResponse>();
    const { error, message } = res.data;

    if (error) {
      throw new Error(message);
    }
  }
}
