interface IDocumentType {
  name: string;
  iconStyle: string;
}

export class DocumentType {
  // TODO: Decode it from `document.proto::Type` if possible
  /** Document types map */
  public static types = new Map<number, IDocumentType>([
    [0, { name: "Document", iconStyle: "" }],
    [1, { name: "Passport", iconStyle: "" }],
    [2, { name: "Internal passport", iconStyle: "" }],
    [3, { name: "Foreign passport", iconStyle: "" }],
    [4, { name: "Identity card", iconStyle: "" }],
    [5, { name: "Driving license", iconStyle: "" }],
    [6, { name: "Hunting license", iconStyle: "" }],
    [7, { name: "Firearms license", iconStyle: "" }],
    [8, { name: "Medical insurance", iconStyle: "" }],
    [9, { name: "Property insurance", iconStyle: "" }],
    [10, { name: "Vehicle insurance", iconStyle: "" }],
    [11, { name: "Personal insurance", iconStyle: "" }],
    [12, { name: "Visa", iconStyle: "" }],
    [13, { name: "Student visa", iconStyle: "" }],
    [14, { name: "Work permit", iconStyle: "" }],
    [15, { name: "Residence permit", iconStyle: "" }],
    [16, { name: "Credit card", iconStyle: "" }],
    [17, { name: "Certificate", iconStyle: "" }],
    [18, { name: "Vaccination Certificate", iconStyle: "" }],
    [19, { name: "Warranty Certificate", iconStyle: "" }],
    [20, { name: "Coupon", iconStyle: "" }],
    [21, { name: "Other", iconStyle: "" }],
  ]);

  public static getName(typeId = 0): string {
    const t = this.types.get(typeId);
    return t ? t.name : "Unknown";
  }

  public static getIconStyle(typeId = 0): string {
    const t = this.types.get(typeId);
    return t ? t.iconStyle : "";
  }
}
