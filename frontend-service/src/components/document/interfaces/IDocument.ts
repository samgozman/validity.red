export interface IDocument {
  ID: string;
  title: string;
  description: string;
  type: string;
  // TODO: Refactor to use a date object
  expiresAt: {
    seconds: number;
  };
}
