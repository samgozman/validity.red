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

export class NotificationService {
  /**
   * Delete notification object
   * @param params
   */
  public static async deleteOne(
    params: NotificationDeletePayload
  ): Promise<void> {
    const payload = JSON.stringify({
      action: "NotificationDelete",
      notification: params,
    });

    const res = await QueryMaker.post(payload);
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
    const payload = JSON.stringify({
      action: "NotificationGetAll",
      notification: {
        documentId,
      },
    });

    const res = await QueryMaker.post<NotificationGetAllResponse>(payload);
    const { error, message, data } = res.data;

    if (error) {
      throw new Error(message);
    }

    return data.notifications;
  }
}
