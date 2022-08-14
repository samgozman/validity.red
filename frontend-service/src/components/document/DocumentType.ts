interface IDocumentType {
  name: string;
  iconStyle: string;
}

export class DocumentType {
  // TODO: Decode it from `document.proto::Type` if possible
  /** Document types map */
  public static types = new Map<number, IDocumentType>([
    [0, { name: "Document", iconStyle: "document-outline" }],
    [1, { name: "Passport", iconStyle: "id-card-outline" }],
    [2, { name: "Internal passport", iconStyle: "id-card-outline" }],
    [3, { name: "Foreign passport", iconStyle: "id-card-outline" }],
    [4, { name: "Identity card", iconStyle: "id-card-outline" }],
    [5, { name: "Driving license", iconStyle: "timer-outline" }],
    [6, { name: "Hunting license", iconStyle: "skull-outline" }],
    [7, { name: "Firearms license", iconStyle: "aperture-outline" }],
    [8, { name: "Medical insurance", iconStyle: "medkit-outline" }],
    [9, { name: "Property insurance", iconStyle: "home-outline" }],
    [10, { name: "Vehicle insurance", iconStyle: "car-sport-outline" }],
    [11, { name: "Personal insurance", iconStyle: "body-outline" }],
    [12, { name: "Visa", iconStyle: "earth-outline" }],
    [13, { name: "Student visa", iconStyle: "school-outline" }],
    [14, { name: "Work permit", iconStyle: "business-outline" }],
    [15, { name: "Residence permit", iconStyle: "ribbon-outline" }],
    [16, { name: "Credit card", iconStyle: "card-outline" }],
    [17, { name: "Certificate", iconStyle: "reader-outline" }],
    [18, { name: "Vaccination Certificate", iconStyle: "qr-code-outline" }],
    [19, { name: "Warranty Certificate", iconStyle: "receipt-outline" }],
    [20, { name: "Coupon", iconStyle: "ticket-outline" }],
    [21, { name: "Other", iconStyle: "document-outline" }],
  ]);

  public static getName(typeId = 0): string {
    const t = this.types.get(typeId);
    return t ? t.name : "Unknown";
  }

  public static getIconStyle(typeId = 0): string {
    const t = this.types.get(typeId);
    return t ? t.iconStyle : "document-outline";
  }
}
