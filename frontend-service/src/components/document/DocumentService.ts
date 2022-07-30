import axios, { type AxiosResponse } from "axios";
import type { IDocument } from "./interfaces/IDocument";

interface IResponse {
  error: boolean;
  message: string;
}

interface DocumentGetAllResponse extends IResponse {
  data: {
    documents: IDocument[];
  };
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

    const res = await this.post<IResponse>(payload);
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

    const res = await this.post<DocumentGetAllResponse>(payload);
    const { error, message, data } = res.data;

    if (error) {
      throw new Error(message);
    }

    return data.documents;
  }

  /**
   * Send POST query to the handler API
   * @param payload Stringified JSON payload
   * @returns
   */
  private static async post<T>(payload: string): Promise<AxiosResponse<T>> {
    return axios.post<T>(`http://localhost:8080/handle`, payload, {
      // To pass Set-Cookie header
      withCredentials: true,
      // Handle 401 error like a normal situation
      validateStatus: (status) =>
        (status >= 200 && status < 300) || status === 401,
    });
  }
}
