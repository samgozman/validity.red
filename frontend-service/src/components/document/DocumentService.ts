import { QueryMaker, type IResponse } from "@/services/QueryMaker";
import type { IDocument } from "./interfaces/IDocument";
import type { INotification } from "./interfaces/INotification";

interface DocumentGetAllResponse extends IResponse {
  data: {
    documents: IDocument[];
  };
}

interface DocumentWithNotifications {
  document: IDocument;
  notifications: INotification[];
}

interface DocumentGetOneResponse extends IResponse {
  data: DocumentWithNotifications;
}

export class DocumentService {
  /**
   * Delete document and all its associated notifications
   * @param documentId
   * @returns
   */
  public static async deleteOne(documentId: string): Promise<void> {
    const payload = JSON.stringify({
      action: "DocumentDelete",
      document: {
        id: documentId,
      },
    });

    const res = await QueryMaker.post(payload);
    const { error, message } = res.data;

    if (error) {
      throw new Error(message);
    }
  }

  /**
   * Get all documents for current user
   * @returns Array of documents
   */
  public static async getAll(): Promise<IDocument[]> {
    const payload = JSON.stringify({
      action: "DocumentGetAll",
    });

    const res = await QueryMaker.post<DocumentGetAllResponse>(payload);
    const { error, message, data } = res.data;

    if (error) {
      throw new Error(message);
    }

    return data.documents;
  }

  public static async getOne(
    documentId: string
  ): Promise<DocumentWithNotifications> {
    const payload = JSON.stringify({
      action: "DocumentGetOne",
      document: {
        id: documentId,
      },
    });

    const res = await QueryMaker.post<DocumentGetOneResponse>(payload);
    const { error, message, data } = res.data;

    if (error) {
      throw new Error(message);
    }

    return data;
  }
}
