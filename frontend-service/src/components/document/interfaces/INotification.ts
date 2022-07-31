export interface INotification {
  ID: string;
  documentID: string;
  // TODO: Refactor to use a date object
  date: {
    seconds: number;
  };
}
