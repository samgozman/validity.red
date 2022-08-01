import { QueryMaker } from "@/services/QueryMaker";

interface NotificationDeletePayload {
  id: string;
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
}
