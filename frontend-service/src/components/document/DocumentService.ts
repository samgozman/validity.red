import { ResponseError } from "@/services/ErrorDecoder";
import { QueryMaker, type IResponse } from "@/services/QueryMaker";
import type { IDocument } from "./interfaces/IDocument";

interface IDocumentGetAllResponse extends IResponse {
  documents: IDocument[];
}

interface IDocumentGetOneResponse extends IResponse {
  document: IDocument;
}

interface IDocumentCreateResponse extends IResponse {
  documentId: string;
}

interface DocumentPayload {
  /** (optional) for document edit */
  id?: string;
  type: number;
  title: string;
  description: string;
  expiresAt: Date;
}

export class DocumentService {
  /**
   * Delete document and all its associated notifications
   * @param documentId
   * @returns
   */
  public static async deleteOne(documentId: string): Promise<void> {
    const res = await new QueryMaker({
      route: `/documents/${documentId}/delete`,
    }).delete();
    const { error, message } = res.data;

    if (error) {
      throw new ResponseError(message);
    }
  }

  /**
   * Get all documents for current user
   * @returns Array of documents
   */
  public static async getAll(): Promise<IDocument[]> {
    const res = await new QueryMaker({
      route: "/documents",
    }).get<IDocumentGetAllResponse>();
    const { error, message, documents } = res.data;

    if (error) {
      throw new ResponseError(message);
    }

    // Short date format
    for (const d of documents) {
      d.expiresAt = this.getDate(d.expiresAt);
    }

    return documents;
  }

  public static async getOne(documentId: string): Promise<IDocument> {
    const res = await new QueryMaker({
      route: `/documents/${documentId}`,
    }).get<IDocumentGetOneResponse>();
    const { error, message, document } = res.data;

    if (error) {
      throw new ResponseError(message);
    }

    // Short date format
    document.expiresAt = this.getDate(document.expiresAt);

    return document;
  }

  public static async createOne(document: DocumentPayload): Promise<string> {
    const payload = JSON.stringify(document);

    const res = await new QueryMaker({
      route: "/documents/create",
      payload,
    }).post<IDocumentCreateResponse>();
    const { error, message, documentId } = res.data;

    if (error) {
      throw new ResponseError(message);
    }

    return documentId;
  }

  public static async updateOne(document: DocumentPayload): Promise<void> {
    const payload = JSON.stringify(document);

    const res = await new QueryMaker({
      route: "/documents/edit",
      payload,
    }).patch<IDocumentCreateResponse>();
    const { error, message } = res.data;

    if (error) {
      throw new ResponseError(message);
    }
  }

  /**
   * Return date in YYYY-MM-DD format
   * @param dateStr
   * @returns
   */
  private static getDate(dateStr: string): string {
    return new Date(dateStr).toISOString().substring(0, 10);
  }
}
