import type { IDocument } from "@/components/document/interfaces/IDocument";
import type { IUsedType } from "./IUsedType";

export interface IDashboardStats {
  totalDocuments: number;
  totalNotifications: number;
  usedTypes: IUsedType[];
  latestDocuments: IDocument[];
}
