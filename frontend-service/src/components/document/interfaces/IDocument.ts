export interface IDocument {
  ID: string;
  title: string;
  description: string;
  type: number;
  // TODO: Refactor to use a date object
  expiresAt: {
    seconds: number;
  };
}
